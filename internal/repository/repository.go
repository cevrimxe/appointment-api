package repository

import (
	"database/sql"
)

type Repositories struct {
	User        UserRepository
	Category    CategoryRepository
	Service     ServiceRepository
	Settings    SettingsRepository
	Device      DeviceRepository
	Specialist  SpecialistRepository
	Appointment AppointmentRepository
	Payment     PaymentRepository
	Contact     ContactRepository
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		User:        NewUserRepository(db),
		Category:    NewCategoryRepository(db),
		Service:     NewServiceRepository(db),
		Settings:    NewSettingsRepository(db),
		Device:      NewDeviceRepository(db),
		Specialist:  NewSpecialistRepository(db),
		Appointment: NewAppointmentRepository(db),
		Payment:     NewPaymentRepository(db),
		Contact:     NewContactRepository(db),
	}
}
