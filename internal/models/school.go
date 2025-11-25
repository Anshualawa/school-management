package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// =====================
//
//	STUDENT MODEL
//
// =====================
type Student struct {
	ID        uuid.UUID `json:"id" gorm:"type:char(36);primaryKey"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:char(36);not null"`
	User      User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	ClassID   uuid.UUID `json:"class_id" gorm:"type:char(36);not null"`
	Class     Class     `gorm:"foreignKey:ClassID"`
	RollNo    int       `json:"roll_no"`
	Address   string    `json:"address"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s *Student) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.New()
	return
}

// =====================
//
//	TEACHER MODEL
//
// =====================
type Teacher struct {
	ID        uuid.UUID `json:"id" gorm:"type:char(36);primaryKey"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:char(36);not null"`
	User      User      `gorm:"foreignKey:UserID"`
	SubjectID uuid.UUID `json:"subject_id" gorm:"type:char(36)"`
	Subject   Subject   `gorm:"foreignKey:SubjectID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t *Teacher) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New()
	return
}

// =====================
//
//	CLASS MODEL
//
// =====================
type Class struct {
	ID        uuid.UUID `json:"id" gorm:"type:char(36);primaryKey"`
	Name      string    `json:"name" gorm:"type:varchar(30);not null"`
	Section   string    `json:"section" gorm:"type:varchar(10)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (c *Class) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New()
	return
}

// =====================
//
//	SUBJECT MODEL
//
// =====================
type Subject struct {
	ID        uuid.UUID `json:"id" gorm:"type:char(36);primaryKey"`
	Name      string    `json:"name" gorm:"type:varchar(100);not null"`
	ClassID   uuid.UUID `json:"class_id" gorm:"type:char(36)"`
	Class     Class     `gorm:"foreignKey:ClassID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s *Subject) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.New()
	return
}

// =====================
//
//	ATTENDANCE MODEL
//
// =====================
type Attendance struct {
	ID        uuid.UUID `json:"id" gorm:"type:char(36);primaryKey"`
	StudentID uuid.UUID `json:"student_id" gorm:"type:char(36)"`
	Student   Student   `gorm:"foreignKey:StudentID"`
	Date      time.Time `json:"date"`
	Status    string    `json:"status" gorm:"type:varchar(10)"` // present/absent
	CreatedAt time.Time
}

func (a *Attendance) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.New()
	return
}

// =====================
//
//	EXAM MODEL
//
// =====================
type Exam struct {
	ID        uuid.UUID `json:"id" gorm:"type:char(36);primaryKey"`
	Name      string    `json:"name"`
	ClassID   uuid.UUID `json:"class_id" gorm:"type:char(36)"`
	Class     Class     `gorm:"foreignKey:ClassID"`
	Date      time.Time `json:"date"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (e *Exam) BeforeCreate(tx *gorm.DB) (err error) {
	e.ID = uuid.New()
	return
}

// =====================
//
//	MARKS MODEL
//
// =====================
type Marks struct {
	ID        uuid.UUID `json:"id" gorm:"type:char(36);primaryKey"`
	StudentID uuid.UUID `json:"student_id" gorm:"type:char(36)"`
	Student   Student   `gorm:"foreignKey:StudentID"`
	ExamID    uuid.UUID `json:"exam_id" gorm:"type:char(36)"`
	Exam      Exam      `gorm:"foreignKey:ExamID"`
	SubjectID uuid.UUID `json:"subject_id" gorm:"type:char(36)"`
	Subject   Subject   `gorm:"foreignKey:SubjectID"`
	Marks     int       `json:"marks"`
	CreatedAt time.Time
}

func (m *Marks) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New()
	return
}

// =====================
//
//	LIBRARY BOOK MODEL
//
// =====================
type Book struct {
	ID        uuid.UUID `json:"id" gorm:"type:char(36);primaryKey"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time
}

func (b *Book) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New()
	return
}

// =====================
//
//	FEES MODEL
//
// =====================
type Fee struct {
	ID        uuid.UUID `json:"id" gorm:"type:char(36);primaryKey"`
	StudentID uuid.UUID `json:"student_id" gorm:"type:char(36)"`
	Student   Student   `gorm:"foreignKey:StudentID"`
	Amount    float64   `json:"amount"`
	Status    string    `json:"status" gorm:"type:varchar(20)"` // paid/unpaid
	Date      time.Time `json:"date"`
	CreatedAt time.Time
}

func (f *Fee) BeforeCreate(tx *gorm.DB) (err error) {
	f.ID = uuid.New()
	return
}

type ClassSubject struct {
	ID        uuid.UUID `json:"id" gorm:"type:char(36);primaryKey"`
	ClassID   uuid.UUID `json:"class_id"`
	SubjectID uuid.UUID `json:"subject_id"`

	Class   Class   `gorm:"foreignKey:ClassID"`
	Subject Subject `gorm:"foreignKey:SubjectID"`
}

func (cs *ClassSubject) BeforeCreate(tx *gorm.DB) (err error) {
	cs.ID = uuid.New()
	return
}

type FeePayment struct {
	ID        uuid.UUID `json:"id" gorm:"type:char(36);primaryKey"`
	StudentID uuid.UUID `json:"student_id"`
	FeeID     uuid.UUID `json:"fee_id"`
	Amount    float64   `json:"amount"`
	PaidAt    time.Time `json:"paid_at"`

	Student Student `gorm:"foreignKey:StudentID"`
	Fee     Fee     `gorm:"foreignKey:FeeID"`
}

func (fp *FeePayment) BeforeCreate(tx *gorm.DB) (err error) {
	fp.ID = uuid.New()
	return
}
