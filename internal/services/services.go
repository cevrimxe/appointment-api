package services

import (
	"appointment-api/internal/config"
	"appointment-api/internal/repository"
	"database/sql"
	"time"
)

type Services struct {
	Auth        AuthService
	Tenant      TenantService
	TenantCache TenantCacheService
	Category    CategoryService
	Service     ServiceService
	Device      DeviceService
	Settings    SettingsService
	User        UserService
	Specialist  SpecialistService
	Appointment AppointmentService
	Payment     PaymentService
	Contact     ContactService
}

func NewServices(repos *repository.Repositories, cfg *config.Config, mainDB *sql.DB) *Services {
	// Create a global user repository for auth (uses main schema)
	globalUserRepo := repository.NewUserRepository(mainDB)

	// Create tenant cache with 5 minute refresh interval
	tenantCache := NewTenantCache(mainDB, 5*time.Minute)

	return &Services{
		Auth:        NewAuthService(globalUserRepo, cfg),
		Tenant:      NewTenantService(mainDB),
		TenantCache: tenantCache,
		Category:    NewCategoryService(repos.Category),
		Service:     NewServiceService(repos.Service, repos.Category),
		Device:      NewDeviceService(repos.Device),
		Settings:    NewSettingsService(repos.Settings, repos.Service),
		User:        NewUserService(repos.User),
		Specialist:  NewSpecialistService(repos.Specialist, repos.Appointment, repos.Settings),
		Appointment: NewAppointmentService(repos.Appointment, repos.Service, repos.Specialist),
		Payment:     NewPaymentService(repos.Payment, repos.Appointment),
		Contact:     NewContactService(repos.Contact),
	}
}
