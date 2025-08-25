<div id="top">

<!-- HEADER STYLE: CONSOLE -->
<div align="center">

```console
  ██   ██████ ██████  ████  ██████ ██   ██ ██████ ██   ██ ██████ ██   ██ ██████          ██   ██████ ██████ 
 ████  ██  ██ ██  ██ ██  ██   ██   ███  ██   ██   ███ ███ ██     ███  ██   ██           ████  ██  ██   ██   
██  ██ ██████ ██████ ██  ██   ██   ██ █ ██   ██   ██ █ ██ ████   ██ █ ██   ██   ██████ ██  ██ ██████   ██   
██████ ██     ██     ██  ██   ██   ██  ███   ██   ██   ██ ██     ██  ███   ██          ██████ ██       ██   
██  ██ ██     ██      ████  ██████ ██   ██   ██   ██   ██ ██████ ██   ██   ██          ██  ██ ██     ██████ 


```

</div>

<!-- BADGES -->
<img src="https://img.shields.io/github/license/cevrimxe/appointment-api?style=flat-square&logo=opensourceinitiative&logoColor=white&color=8a2be2" alt="license">
<img src="https://img.shields.io/github/last-commit/cevrimxe/appointment-api?style=flat-square&logo=git&logoColor=white&color=8a2be2" alt="last-commit">
<img src="https://img.shields.io/github/languages/top/cevrimxe/appointment-api?style=flat-square&color=8a2be2" alt="repo-top-language">
<img src="https://img.shields.io/github/languages/count/cevrimxe/appointment-api?style=flat-square&color=8a2be2" alt="repo-language-count">

<em>Built with the tools and technologies:</em>

<img src="https://img.shields.io/badge/Go-00ADD8.svg?style=flat-square&logo=Go&logoColor=white" alt="Go">
<img src="https://img.shields.io/badge/Gin-008ECF.svg?style=flat-square&logo=Gin&logoColor=white" alt="Gin">
<img src="https://img.shields.io/badge/YAML-CB171E.svg?style=flat-square&logo=YAML&logoColor=white" alt="YAML">

</div>
<br>

## 💧 Table of Contents

<details>
<summary>Table of Contents</summary>

- [💧 Table of Contents](#-table-of-contents)
- [🌊 Overview](#-overview)
- [💦 Features](#-features)
- [🔵 Project Structure](#-project-structure)
    - [🔷 Project Index](#-project-index)
- [💠 Getting Started](#-getting-started)
    - [🅿️ Prerequisites](#-prerequisites)
    - [🌀 Installation](#-installation)
    - [🔹 Usage](#-usage)
    - [❄ ️ Testing](#-testing)
- [🧊 Roadmap](#-roadmap)
- [⚪ Contributing](#-contributing)
- [⬜ License](#-license)
- [✨ Acknowledgments](#-acknowledgments)

</details>

---

## 🌊 Overview

The Appointment API is a comprehensive REST API service built with Go and Gin framework, designed to manage appointments, users, and services in a multi-tenant environment. This system provides a robust backend for scheduling applications, supporting features like user authentication, appointment management, service categorization, and payment processing. It's built with scalability and security in mind, using modern Go practices and a clean architecture approach.


---

## 💦 Features

- **🔐 Multi-tenant Architecture**: Support for multiple business entities with isolated data
- **👥 User Management**: Complete user authentication and authorization system
- **📅 Appointment Scheduling**: Flexible appointment booking and management
- **💼 Service Management**: Categorization and management of services
- **💳 Payment Processing**: Integrated payment handling system
- **📱 Device Management**: Support for multiple devices and notifications
- **⚙️ Settings Management**: Customizable system settings per tenant
- **📊 Reporting**: Built-in reporting capabilities
- **🔄 CORS Support**: Cross-Origin Resource Sharing enabled
- **📝 Logging**: Comprehensive logging system
- **🔒 Security**: JWT-based authentication and middleware protection

---

## 🔵 Project Structure

```sh
└── appointment-api/
    ├── README.md
    ├── admin-endpoints.md
    ├── adminedit.md
    ├── appointment-api
    ├── cmd
    │   └── server
    ├── config.env
    ├── endpoints.md
    ├── go.mod
    ├── go.sum
    ├── internal
    │   ├── api
    │   ├── config
    │   ├── middleware
    │   ├── models
    │   ├── repository
    │   └── services
    ├── migrations
    │   ├── 001_public_schema.sql
    │   └── 005_complete_tenant_schema.sql
    └── tests
        ├── README.md
        └── test_runner.go
```

### 🔷 Project Index

<details open>
	<summary><b><code>APPOINTMENT-API/</code></b></summary>
	<!-- __root__ Submodule -->
	<details>
		<summary><b>__root__</b></summary>
		<blockquote>
			<div class='directory-path' style='padding: 8px 0; color: #666;'>
				<code><b>⦿ __root__</b></code>
			<table style='width: 100%; border-collapse: collapse;'>
			<thead>
				<tr style='background-color: #f8f9fa;'>
					<th style='width: 30%; text-align: left; padding: 8px;'>File Name</th>
					<th style='text-align: left; padding: 8px;'>Summary</th>
				</tr>
			</thead>
				<tr style='border-bottom: 1px solid #eee;'>
					<td style='padding: 8px;'><b><a href='https://github.com/cevrimxe/appointment-api/blob/master/go.sum'>go.sum</a></b></td>
					<td style='padding: 8px;'>Go module dependency checksums</td>
				</tr>
				<tr style='border-bottom: 1px solid #eee;'>
					<td style='padding: 8px;'><b><a href='https://github.com/cevrimxe/appointment-api/blob/master/appointment-api'>appointment-api</a></b></td>
					<td style='padding: 8px;'>Main executable binary file</td>
				</tr>
				<tr style='border-bottom: 1px solid #eee;'>
					<td style='padding: 8px;'><b><a href='https://github.com/cevrimxe/appointment-api/blob/master/go.mod'>go.mod</a></b></td>
					<td style='padding: 8px;'>Go module definition and dependency management</td>
				</tr>
			</table>
		</blockquote>
	</details>
	<!-- cmd Submodule -->
	<details>
		<summary><b>cmd</b></summary>
		<blockquote>
			<div class='directory-path' style='padding: 8px 0; color: #666;'>
				<code><b>⦿ cmd</b></code>
			<!-- server Submodule -->
			<details>
				<summary><b>server</b></summary>
				<blockquote>
					<div class='directory-path' style='padding: 8px 0; color: #666;'>
						<code><b>⦿ cmd.server</b></code>
					<table style='width: 100%; border-collapse: collapse;'>
					<thead>
						<tr style='background-color: #f8f9fa;'>
							<th style='width: 30%; text-align: left; padding: 8px;'>File Name</th>
							<th style='text-align: left; padding: 8px;'>Summary</th>
						</tr>
					</thead>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/cevrimxe/appointment-api/blob/master/cmd/server/main.go'>main.go</a></b></td>
							<td style='padding: 8px;'>Application entry point and server initialization</td>
						</tr>
					</table>
				</blockquote>
			</details>
		</blockquote>
	</details>
	<!-- migrations Submodule -->
	<details>
		<summary><b>migrations</b></summary>
		<blockquote>
			<div class='directory-path' style='padding: 8px 0; color: #666;'>
				<code><b>⦿ migrations</b></code>
			<table style='width: 100%; border-collapse: collapse;'>
			<thead>
				<tr style='background-color: #f8f9fa;'>
					<th style='width: 30%; text-align: left; padding: 8px;'>File Name</th>
					<th style='text-align: left; padding: 8px;'>Summary</th>
				</tr>
			</thead>
				<tr style='border-bottom: 1px solid #eee;'>
					<td style='padding: 8px;'><b><a href='https://github.com/cevrimxe/appointment-api/blob/master/migrations/001_public_schema.sql'>001_public_schema.sql</a></b></td>
					<td style='padding: 8px;'>Initial database schema setup</td>
				</tr>
				<tr style='border-bottom: 1px solid #eee;'>
					<td style='padding: 8px;'><b><a href='https://github.com/cevrimxe/appointment-api/blob/master/migrations/005_complete_tenant_schema.sql'>005_complete_tenant_schema.sql</a></b></td>
					<td style='padding: 8px;'>Multi-tenant schema implementation</td>
				</tr>
			</table>
		</blockquote>
	</details>
	<!-- internal Submodule -->
	<details>
		<summary><b>internal</b></summary>
		<blockquote>
			<div class='directory-path' style='padding: 8px 0; color: #666;'>
				<code><b>⦿ internal</b></code>
			<!-- api Submodule -->
			<details>
				<summary><b>api</b></summary>
				<blockquote>
					<div class='directory-path' style='padding: 8px 0; color: #666;'>
						<code><b>⦿ internal.api</b></code>
					<table style='width: 100%; border-collapse: collapse;'>
					<thead>
						<tr style='background-color: #f8f9fa;'>
							<th style='width: 30%; text-align: left; padding: 8px;'>File Name</th>
							<th style='text-align: left; padding: 8px;'>Summary</th>
						</tr>
					</thead>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/cevrimxe/appointment-api/blob/master/internal/api/auth_handler.go'>auth_handler.go</a></b></td>
							<td style='padding: 8px;'>Authentication and authorization request handlers</td>
						</tr>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/cevrimxe/appointment-api/blob/master/internal/api/handlers.go'>handlers.go</a></b></td>
							<td style='padding: 8px;'>Common API request handlers and utilities</td>
						</tr>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/cevrimxe/appointment-api/blob/master/internal/api/admin_handler.go'>admin_handler.go</a></b></td>
							<td style='padding: 8px;'>Administrative endpoints and operations</td>
						</tr>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/cevrimxe/appointment-api/blob/master/internal/api/public_handler.go'>public_handler.go</a></b></td>
							<td style='padding: 8px;'>Public API endpoints and handlers</td>
						</tr>
					</table>
				</blockquote>
			</details>
			<!-- config Submodule -->
			<details>
				<summary><b>config</b></summary>
				<blockquote>
					<div class='directory-path' style='padding: 8px 0; color: #666;'>
						<code><b>⦿ internal.config</b></code>
					<table style='width: 100%; border-collapse: collapse;'>
					<thead>
						<tr style='background-color: #f8f9fa;'>
							<th style='width: 30%; text-align: left; padding: 8px;'>File Name</th>
							<th style='text-align: left; padding: 8px;'>Summary</th>
						</tr>
					</thead>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/cevrimxe/appointment-api/blob/master/internal/config/config.go'>config.go</a></b></td>
							<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
						</tr>
					</table>
				</blockquote>
			</details>
			<!-- repository Submodule -->
			<details>
				<summary><b>repository</b></summary>
				<blockquote>
					<div class='directory-path' style='padding: 8px 0; color: #666;'>
						<code><b>⦿ internal.repository</b></code>
					<table style='width: 100%; border-collapse: collapse;'>
					<thead>
						<tr style='background-color: #f8f9fa;'>
							<th style='width: 30%; text-align: left; padding: 8px;'>File Name</th>
							<th style='text-align: left; padding: 8px;'>Summary</th>
						</tr>
					</thead>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/cevrimxe/appointment-api/blob/master/internal/repository/payment_repository.go'>payment_repository.go</a></b></td>
							<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
						</tr>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/cevrimxe/appointment-api/blob/master/internal/repository/contact_repository.go'>contact_repository.go</a></b></td>
							<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
						</tr>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/cevrimxe/appointment-api/blob/master/internal/repository/device_repository.go'>device_repository.go</a></b></td>
							<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
						</tr>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/cevrimxe/appointment-api/blob/master/internal/repository/user_repository.go'>user_repository.go</a></b></td>
							<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
						</tr>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/cevrimxe/appointment-api/blob/master/internal/repository/repository.go'>repository.go</a></b></td>
							<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
						</tr>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/cevrimxe/appointment-api/blob/master/internal/repository/category_repository.go'>category_repository.go</a></b></td>
							<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
						</tr>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/cevrimxe/appointment-api/blob/master/internal/repository/settings_repository.go'>settings_repository.go</a></b></td>
							<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
						</tr>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/cevrimxe/appointment-api/blob/master/internal/repository/service_repository.go'>service_repository.go</a></b></td>
							<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
						</tr>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/cevrimxe/appointment-api/blob/master/internal/repository/appointment_repository.go'>appointment_repository.go</a></b></td>
							<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
						</tr>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/cevrimxe/appointment-api/blob/master/internal/repository/specialist_repository.go'>specialist_repository.go</a></b></td>
							<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
						</tr>
					</table>
				</blockquote>
			</details>
			<!-- middleware Submodule -->
			<details>
				<summary><b>middleware</b></summary>
				<blockquote>
					<div class='directory-path' style='padding: 8px 0; color: #666;'>
						<code><b>⦿ internal.middleware</b></code>
					<table style='width: 100%; border-collapse: collapse;'>
					<thead>
						<tr style='background-color: #f8f9fa;'>
							<th style='width: 30%; text-align: left; padding: 8px;'>File Name</th>
							<th style='text-align: left; padding: 8px;'>Summary</th>
						</tr>
					</thead>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/cevrimxe/appointment-api/blob/master/internal/middleware/auth.go'>auth.go</a></b></td>
							<td style='padding: 8px;'>Authentication middleware for request validation</td>
						</tr>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/cevrimxe/appointment-api/blob/master/internal/middleware/tenant.go'>tenant.go</a></b></td>
							<td style='padding: 8px;'>Multi-tenant request processing middleware</td>
						</tr>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/cevrimxe/appointment-api/blob/master/internal/middleware/logging.go'>logging.go</a></b></td>
							<td style='padding: 8px;'>Request logging and monitoring middleware</td>
						</tr>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/cevrimxe/appointment-api/blob/master/internal/middleware/cors.go'>cors.go</a></b></td>
							<td style='padding: 8px;'>Cross-Origin Resource Sharing configuration</td>
						</tr>
					</table>
				</blockquote>
			</details>
			<!-- services Submodule -->
			<details>
				<summary><b>services</b></summary>
				<blockquote>
					<div class='directory-path' style='padding: 8px 0; color: #666;'>
						<code><b>⦿ internal.services</b></code>
					<table style='width: 100%; border-collapse: collapse;'>
					<thead>
						<tr style='background-color: #f8f9fa;'>
							<th style='width: 30%; text-align: left; padding: 8px;'>File Name</th>
							<th style='text-align: left; padding: 8px;'>Summary</th>
						</tr>
					</thead>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/cevrimxe/appointment-api/blob/master/internal/services/tenant_service.go'>tenant_service.go</a></b></td>
							<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
						</tr>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/cevrimxe/appointment-api/blob/master/internal/services/contact_service.go'>contact_service.go</a></b></td>
							<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
						</tr>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/cevrimxe/appointment-api/blob/master/internal/services/category_service.go'>category_service.go</a></b></td>
							<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
						</tr>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/cevrimxe/appointment-api/blob/master/internal/services/specialist_service.go'>specialist_service.go</a></b></td>
							<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
						</tr>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/cevrimxe/appointment-api/blob/master/internal/services/payment_service.go'>payment_service.go</a></b></td>
							<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
						</tr>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/cevrimxe/appointment-api/blob/master/internal/services/device_service.go'>device_service.go</a></b></td>
							<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
						</tr>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/cevrimxe/appointment-api/blob/master/internal/services/tenant_cache.go'>tenant_cache.go</a></b></td>
							<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
						</tr>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/cevrimxe/appointment-api/blob/master/internal/services/user_service.go'>user_service.go</a></b></td>
							<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
						</tr>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/cevrimxe/appointment-api/blob/master/internal/services/service_service.go'>service_service.go</a></b></td>
							<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
						</tr>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/cevrimxe/appointment-api/blob/master/internal/services/upload_service.go'>upload_service.go</a></b></td>
							<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
						</tr>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/cevrimxe/appointment-api/blob/master/internal/services/appointment_service.go'>appointment_service.go</a></b></td>
							<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
						</tr>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/cevrimxe/appointment-api/blob/master/internal/services/services.go'>services.go</a></b></td>
							<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
						</tr>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/cevrimxe/appointment-api/blob/master/internal/services/settings_service.go'>settings_service.go</a></b></td>
							<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
						</tr>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/cevrimxe/appointment-api/blob/master/internal/services/auth_service.go'>auth_service.go</a></b></td>
							<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
						</tr>
					</table>
				</blockquote>
			</details>
			<!-- models Submodule -->
			<details>
				<summary><b>models</b></summary>
				<blockquote>
					<div class='directory-path' style='padding: 8px 0; color: #666;'>
						<code><b>⦿ internal.models</b></code>
					<table style='width: 100%; border-collapse: collapse;'>
					<thead>
						<tr style='background-color: #f8f9fa;'>
							<th style='width: 30%; text-align: left; padding: 8px;'>File Name</th>
							<th style='text-align: left; padding: 8px;'>Summary</th>
						</tr>
					</thead>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/cevrimxe/appointment-api/blob/master/internal/models/settings.go'>settings.go</a></b></td>
							<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
						</tr>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/cevrimxe/appointment-api/blob/master/internal/models/payment.go'>payment.go</a></b></td>
							<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
						</tr>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/cevrimxe/appointment-api/blob/master/internal/models/device.go'>device.go</a></b></td>
							<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
						</tr>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/cevrimxe/appointment-api/blob/master/internal/models/category.go'>category.go</a></b></td>
							<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
						</tr>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/cevrimxe/appointment-api/blob/master/internal/models/user.go'>user.go</a></b></td>
							<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
						</tr>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/cevrimxe/appointment-api/blob/master/internal/models/report.go'>report.go</a></b></td>
							<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
						</tr>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/cevrimxe/appointment-api/blob/master/internal/models/contact.go'>contact.go</a></b></td>
							<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
						</tr>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/cevrimxe/appointment-api/blob/master/internal/models/tenant.go'>tenant.go</a></b></td>
							<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
						</tr>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/cevrimxe/appointment-api/blob/master/internal/models/appointment.go'>appointment.go</a></b></td>
							<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
						</tr>
					</table>
				</blockquote>
			</details>
		</blockquote>
	</details>
</details>

---

## 💠 Getting Started

### 🅿️ Prerequisites

This project requires the following dependencies:

- **Programming Language:** Go
- **Package Manager:** Go modules

### 🌀 Installation

Build appointment-api from the source and intsall dependencies:

1. **Clone the repository:**

    ```sh
    ❯ git clone https://github.com/cevrimxe/appointment-api
    ```

2. **Navigate to the project directory:**

    ```sh
    ❯ cd appointment-api
    ```

3. **Install the dependencies:**

    ```sh
    ❯ go mod download
    ```

4. **Build the project:**

    ```sh
    ❯ go build ./cmd/server
    ```

### 🔹 Usage

Run the project with:

**Using [go modules](https://golang.org/):**
```sh
go run {entrypoint}
```

### ❄️ Testing

Appointment-api uses Go's built-in testing package with testify assertions for enhanced testing capabilities. Run the test suite with:

**Using [go modules](https://golang.org/):**
```sh
go test ./...
```

---

## 🧊 Roadmap

- [X] **`v1.0.0`**: <strike>Core API implementation with basic appointment and user management</strike>
- [ ] **`v1.1.0`**: Advanced scheduling features with recurring appointments
- [ ] **`v1.2.0`**: Enhanced reporting and analytics
- [ ] **`v1.3.0`**: Mobile app integration with push notifications
- [ ] **`v2.0.0`**: Real-time scheduling and instant notifications

---

## ⚪ Contributing

- **💬 [Join the Discussions](https://github.com/cevrimxe/appointment-api/discussions)**: Share your insights, provide feedback, or ask questions.
- **🐛 [Report Issues](https://github.com/cevrimxe/appointment-api/issues)**: Submit bugs found or log feature requests for the `appointment-api` project.
- **💡 [Submit Pull Requests](https://github.com/cevrimxe/appointment-api/blob/main/CONTRIBUTING.md)**: Review open PRs, and submit your own PRs.

<details closed>
<summary>Contributing Guidelines</summary>

1. **Fork the Repository**: Start by forking the project repository to your github account.
2. **Clone Locally**: Clone the forked repository to your local machine using a git client.
   ```sh
   git clone https://github.com/cevrimxe/appointment-api
   ```
3. **Create a New Branch**: Always work on a new branch, giving it a descriptive name.
   ```sh
   git checkout -b new-feature-x
   ```
4. **Make Your Changes**: Develop and test your changes locally.
5. **Commit Your Changes**: Commit with a clear message describing your updates.
   ```sh
   git commit -m 'Implemented new feature x.'
   ```
6. **Push to github**: Push the changes to your forked repository.
   ```sh
   git push origin new-feature-x
   ```
7. **Submit a Pull Request**: Create a PR against the original project repository. Clearly describe the changes and their motivations.
8. **Review**: Once your PR is reviewed and approved, it will be merged into the main branch. Congratulations on your contribution!
</details>

<details closed>
<summary>Contributor Graph</summary>
<br>
<p align="left">
   <a href="https://github.com{/cevrimxe/appointment-api/}graphs/contributors">
      <img src="https://contrib.rocks/image?repo=cevrimxe/appointment-api">
   </a>
</p>
</details>

---

## ⬜ License

Appointment-api is protected under the [LICENSE](https://choosealicense.com/licenses) License. For more details, refer to the [LICENSE](https://choosealicense.com/licenses/) file.

---

## ✨ Acknowledgments

- Thanks to the Go community for excellent documentation and support
- [Gin Web Framework](https://gin-gonic.com/) for providing a robust foundation
- [JWT-Go](https://github.com/golang-jwt/jwt) for secure authentication
- All contributors who have helped shape and improve this project
- Special thanks to the open source community for inspiration and shared knowledge

<div align="right">

[![][back-to-top]](#top)

</div>


[back-to-top]: https://img.shields.io/badge/-BACK_TO_TOP-151515?style=flat-square


---
