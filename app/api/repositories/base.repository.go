package repositories

import "github.com/jinzhu/gorm"

// BaseRepository ...
type BaseRepository interface {
	FindAll(v interface{}) (err error)
	FindOneByID(v interface{}, id int64) (err error)
	Create(v interface{}) error
	Delete(v interface{}) error
}

// NewBaseRepository ...
func NewBaseRepository(db *gorm.DB) BaseRepository {
	return &baseRepository{
		db: db,
	}
}

type baseRepository struct {
	db *gorm.DB
}

type (
	// DBFunc type which accept *gorm.DB and return error
	DBFunc func(tx *gorm.DB) error
)

// Create
// Helper function to insert gorm model to database by using 'WithinTransaction'
func (r *baseRepository) Create(v interface{}) error {
	return r.WithinTransaction(func(tx *gorm.DB) (err error) {
		// check new object
		if !r.db.NewRecord(v) {
			return err
		}
		if err = tx.Create(v).Error; err != nil {
			tx.Rollback() // rollback
			return err
		}
		return err
	})
}

// // Save
// // Helper function to save gorm model to database by using 'WithinTransaction'
// func (r *baseRepository) Save(v interface{}) error {
// 	return WithinTransaction(func(tx *gorm.DB) (err error) {
// 		// check new object
// 		if cgorm.DBManager().NewRecord(v) {
// 			return err
// 		}
// 		if err = tx.Save(v).Error; err != nil {
// 			tx.Rollback()                                                                                                       // rollback
// 			return err
// 		}
// 		return err
// 	})
// }

// Delete
// Helper function to save gorm model to database by using 'WithinTransaction'
func (r *baseRepository) Delete(v interface{}) error {
	return r.WithinTransaction(func(tx *gorm.DB) (err error) {
		// check new object
		if err = tx.Delete(v).Error; err != nil {
			tx.Rollback() // rollback
			return err
		}
		return err
	})
}

// FindOneByID
// Helper function to find a record by using 'WithinTransaction'
func (r *baseRepository) FindOneByID(v interface{}, id int64) (err error) {
	return r.WithinTransaction(func(tx *gorm.DB) error {
		if err = tx.Last(v, id).Error; err != nil {
			tx.Rollback() // rollback db transaction
			return err
		}
		return err
	})
}

// FindAll
// Helper function to find records by using 'WithinTransaction'
func (r *baseRepository) FindAll(v interface{}) (err error) {
	return r.WithinTransaction(func(tx *gorm.DB) error {
		if err = tx.Find(v).Error; err != nil {
			tx.Rollback() // rollback db transaction
			return err
		}
		return err
	})
}

// // FindOneByQuery
// // Helper function to find a record by using 'WithinTransaction'
// func (r *baseRepository) FindOneByQuery(v interface{}, params map[string]interface{}) (err error) {
// 	return WithinTransaction(func(tx *gorm.DB) error {
// 		if err = tx.Where(params).Last(v).Error; err != nil {
// 			tx.Rollback() // rollback db transaction
// 			return err
// 		}
// 		return err
// 	})
// }

// FindByQuery
// Helper function to find records by using 'WithinTransaction'
func (r *baseRepository) FindByQuery(v interface{}, params map[string]interface{}) (err error) {
	return r.WithinTransaction(func(tx *gorm.DB) error {
		if err = tx.Where(params).Find(v).Error; err != nil {
			tx.Rollback() // rollback db transaction
			return err
		}
		return err
	})
}

// WithinTransaction
// accept DBFunc as parameter
// call DBFunc function within transaction begin, and commit and return error from DBFunc
func (r *baseRepository) WithinTransaction(fn DBFunc) (err error) {
	tx := r.db.Begin() // start db transaction
	defer tx.Commit()
	err = fn(tx)
	// close db transaction
	return err
}
