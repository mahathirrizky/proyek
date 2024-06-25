// file: adminproyek/formatter.go

package adminproyek

import (
    "time"
)

type AdminProyekFormatter struct {
    IDAdminProyek int       `json:"id_admin_proyek"`
    IDUser        int       `json:"id_user"`
    IDProyek      int       `json:"id_proyek"`
    CreatedAt     time.Time `json:"created_at"`
    UpdatedAt     time.Time `json:"updated_at"`
}

func FormatAdminProyek(adminProyek AdminProyekTable) AdminProyekFormatter {
    return AdminProyekFormatter{
        IDAdminProyek: adminProyek.IdAdminProyek,
        IDUser:        adminProyek.IdUser,
        IDProyek:      adminProyek.IdProyek,
        CreatedAt:     adminProyek.CreatedAt,
        UpdatedAt:     adminProyek.UpdatedAt,
    }
}

func FormatAdminProyeks(adminProyeks []AdminProyekTable) []AdminProyekFormatter {
    var adminProyekFormatters []AdminProyekFormatter
    for _, adminProyek := range adminProyeks {
        adminProyekFormatters = append(adminProyekFormatters, FormatAdminProyek(adminProyek))
    }
    return adminProyekFormatters
}
