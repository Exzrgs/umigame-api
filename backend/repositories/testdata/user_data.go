package testdata

import (
	"umigame-api/models"
)

var (
	defaultName         = "test"
	defaultEmail        = "isujiv@peltul.lk"
	defaultPasswordHash = "ee26b0dd4af7e749aa1a8ee3c10ae9923f618980772e473f8819a5d4940e0db27ac185f8a0e1d5f84f88bc887fd67b143732c304cc5fa9ad8e6f57f50028a8ff" // "test"
	defaultUUID         = "d04e991f-f272-59ff-aadd-e312c7d73f89"
	defaultIsValid      = true

	data1Email = "ajalgun@fe.gy"
	data1UUID  = "7c19a998-eac3-55be-827e-05809d977ce8"
)

var SelectUser_Basic = []models.User{
	{
		Name:         defaultName,
		Email:        data1Email,
		PasswordHash: defaultPasswordHash,
		UUID:         data1UUID,
		IsValid:      defaultIsValid,
	},
}

var InsertUser_Basic = []models.User{
	{
		Name:         defaultName,
		Email:        defaultEmail,
		PasswordHash: defaultPasswordHash,
		UUID:         defaultUUID,
		IsValid:      false,
	},
}

var InsertUser_Char = []models.User{
	{
		Name:         "テスト",
		Email:        defaultEmail,
		PasswordHash: defaultPasswordHash,
		UUID:         defaultUUID,
		IsValid:      false,
	},
	{
		Name:         "てすと",
		Email:        defaultEmail,
		PasswordHash: defaultPasswordHash,
		UUID:         defaultUUID,
		IsValid:      false,
	},
	{
		Name:         "点検",
		Email:        defaultEmail,
		PasswordHash: defaultPasswordHash,
		UUID:         defaultUUID,
		IsValid:      false,
	},
	{
		Name:         "ｱｲｳｴｵ",
		Email:        defaultEmail,
		PasswordHash: defaultPasswordHash,
		UUID:         defaultUUID,
		IsValid:      false,
	},
	{
		Name:         `!"#$%&'()=-~^：:；＠」「・。、￥＾ー|\[{]}]*/?.>,<\_`,
		Email:        defaultEmail,
		PasswordHash: defaultPasswordHash,
		UUID:         defaultUUID,
		IsValid:      false,
	},
	{
		Name:         "　　　　　    ",
		Email:        defaultEmail,
		PasswordHash: defaultPasswordHash,
		UUID:         defaultUUID,
		IsValid:      false,
	},
	{
		Name:         "✲𝑩𝑰𝑮 𝑳𝑶𝑽𝑬😎𖧡",
		Email:        defaultEmail,
		PasswordHash: defaultPasswordHash,
		UUID:         defaultUUID,
		IsValid:      false,
	},
}

var InsertUser_TooLong = []models.User{
	{
		Name:         "tttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttt",
		Email:        defaultEmail,
		PasswordHash: defaultPasswordHash,
		UUID:         defaultUUID,
		IsValid:      defaultIsValid,
	},
	{
		Name:         defaultName,
		Email:        "tttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttt",
		PasswordHash: defaultPasswordHash,
		UUID:         defaultUUID,
		IsValid:      defaultIsValid,
	},
	{
		Name:         defaultName,
		Email:        defaultEmail,
		PasswordHash: defaultPasswordHash,
		UUID:         "tttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttt",
		IsValid:      defaultIsValid,
	},
}

// //////////////////////////////////////////////////////

var UpdateValid_Basic = []models.User{
	{
		Name:         defaultName,
		Email:        defaultEmail,
		PasswordHash: defaultPasswordHash,
		UUID:         defaultUUID,
		IsValid:      false,
	},
}

// //////////////////////////////////////////////////////

var GetAuthInfo_Basic = []models.User{
	{
		Name:         defaultName,
		Email:        data1Email,
		PasswordHash: defaultPasswordHash,
		UUID:         data1UUID,
		IsValid:      defaultIsValid,
	},
}

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
