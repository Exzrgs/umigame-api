package testdata

// import (
// 	"umigame-api/models"
// )

// var (
// 	defaultName         = "test"
// 	defaultEmail        = "ajalgun@fe.gy"
// 	defaultPasswordHash = "ee26b0dd4af7e749aa1a8ee3c10ae9923f618980772e473f8819a5d4940e0db27ac185f8a0e1d5f84f88bc887fd67b143732c304cc5fa9ad8e6f57f50028a8ff"
// 	defauktUuid         = "491d233f-8d0f-5090-88e3-03ac8a560a5a"
// 	defaultIsValid      = true
// )

// var InsertUser_Basic = []models.User{
// 	{
// 		Name:         defaultName,
// 		Email:        defaultEmail,
// 		PasswordHash: defaultPasswordHash,
// 		UUID:         defauktUuid,
// 		IsValid:      false,
// 	},
// }

// var InsertUser_NoData = []models.User{
// 	{
// 		// Name:         defaultName,
// 		Email:        defaultEmail,
// 		PasswordHash: defaultPasswordHash,
// 		UUID:         defauktUuid,
// 		IsValid:      false,
// 	},
// 	{
// 		Name: defaultName,
// 		// Email:        defaultEmail,
// 		PasswordHash: defaultPasswordHash,
// 		UUID:         defauktUuid,
// 		IsValid:      false,
// 	},
// 	{
// 		Name:  defaultName,
// 		Email: defaultEmail,
// 		// PasswordHash: defaultPasswordHash,
// 		UUID:    defauktUuid,
// 		IsValid: false,
// 	},
// 	{
// 		Name:         defaultName,
// 		Email:        defaultEmail,
// 		PasswordHash: defaultPasswordHash,
// 		// UUID:         defauktUuid,
// 		IsValid: false,
// 	},
// 	{
// 		Name:         defaultName,
// 		Email:        defaultEmail,
// 		PasswordHash: defaultPasswordHash,
// 		UUID:         defauktUuid,
// 		// IsValid:      false,
// 	},
// }

// var InsertUser_Char = []models.User{
// 	{
// 		Name:         "テスト",
// 		Email:        defaultEmail,
// 		PasswordHash: defaultPasswordHash,
// 		UUID:         defauktUuid,
// 		IsValid:      false,
// 	},
// 	{
// 		Name:         "てすと",
// 		Email:        defaultEmail,
// 		PasswordHash: defaultPasswordHash,
// 		UUID:         defauktUuid,
// 		IsValid:      false,
// 	},
// 	{
// 		Name:         "点検",
// 		Email:        defaultEmail,
// 		PasswordHash: defaultPasswordHash,
// 		UUID:         defauktUuid,
// 		IsValid:      false,
// 	},
// 	{
// 		Name:         "ｱｲｳｴｵ",
// 		Email:        defaultEmail,
// 		PasswordHash: defaultPasswordHash,
// 		UUID:         defauktUuid,
// 		IsValid:      false,
// 	},
// 	{
// 		Name:         `!"#$%&'()=-~^：:；＠」「・。、￥＾ー|\[{]}]*/?.>,<\_`,
// 		Email:        defaultEmail,
// 		PasswordHash: defaultPasswordHash,
// 		UUID:         defauktUuid,
// 		IsValid:      false,
// 	},
// 	{
// 		Name:         "　　　　　    ",
// 		Email:        defaultEmail,
// 		PasswordHash: defaultPasswordHash,
// 		UUID:         defauktUuid,
// 		IsValid:      false,
// 	},
// 	{
// 		Name:         "✲𝑩𝑰𝑮 𝑳𝑶𝑽𝑬😎𖧡",
// 		Email:        defaultEmail,
// 		PasswordHash: defaultPasswordHash,
// 		UUID:         defauktUuid,
// 		IsValid:      false,
// 	},
// }

// var InsertUser_TooLong = []models.User{
// 	{
// 		Name:         "tttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttt",
// 		Email:        defaultEmail,
// 		PasswordHash: defaultPasswordHash,
// 		UUID:         defauktUuid,
// 		IsValid:      defaultIsValid,
// 	},
// 	{
// 		Name:         defaultName,
// 		Email:        "tttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttt",
// 		PasswordHash: defaultPasswordHash,
// 		UUID:         defauktUuid,
// 		IsValid:      defaultIsValid,
// 	},
// 	{
// 		Name:         defaultName,
// 		Email:        defaultEmail,
// 		PasswordHash: "tttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttt",
// 		UUID:         defauktUuid,
// 		IsValid:      defaultIsValid,
// 	},
// 	{
// 		Name:         defaultName,
// 		Email:        defaultEmail,
// 		PasswordHash: defaultPasswordHash,
// 		UUID:         "tttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttt",
// 		IsValid:      defaultIsValid,
// 	},
// }

// //////////////////////////////////////////////////////

// var UpdateValid_Basic = []models.User{
// 	{
// 		Name:         defaultName,
// 		Email:        defaultEmail,
// 		PasswordHash: defaultPasswordHash,
// 		UUID:         defauktUuid,
// 		IsValid:      true,
// 	},
// }

// //////////////////////////////////////////////////////

// var GetAuthInfo_Basic = []models.User{
// 	{
// 		Name:         defaultName,
// 		PasswordHash: defaultPasswordHash,
// 		UUID:         defauktUuid,
// 		IsValid:      defaultIsValid,
// 	},
// }

// //////////////////////////////////////////////////////

// var DeleteUser = []models.User{
// 	{},
// }

// //////////////////////////////////////////////////////

// var UpdateEmail = []models.User{
// 	{
// 		Name:         defaultName,
// 		Email:        "as@amuduhi.qa",
// 		PasswordHash: defaultPasswordHash,
// 		UUID:         defauktUuid,
// 		IsValid:      false,
// 	},
// }
