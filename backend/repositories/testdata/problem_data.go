package testdata

import (
	"time"
	"umigame-api/models"
)

var testDate = time.Date(2023, 1, 28, 12, 0, 0, 0, time.UTC)

///////////////////////////////////////////////////////////

/*
TODO:CreatedAtを入れないとテストできない。
テストデータに入れないにしても、テスト関数では何らかの方法でしっかりチェックしないといけない
*/
var SelectProblemList_Basic = []models.Problem{
	{
		ID:        1,
		Title:     "test1",
		Author:    "test1",
		Statement: "test1",
		CreatedAt: &testDate,
	},
	{
		ID:        2,
		Title:     "test2",
		Author:    "test2",
		Statement: "test2",
		CreatedAt: &testDate,
	},
	{
		ID:        3,
		Title:     "test3",
		Author:    "test3",
		Statement: "test3",
		CreatedAt: &testDate,
	},
}

///////////////////////////////////////////////////////////

var SelectProblem_Basic = []models.Problem{
	{
		ID:           1,
		Title:        "test1",
		Statement:    "test1",
		Answer:       "test1",
		Author:       "test1",
		Reference:    "test1",
		ReferenceURL: "test1",
		CreatedAt:    &testDate,
	},
}

///////////////////////////////////////////////////////////

var InsertProblem_Basic = []models.Problem{
	{
		Title:        "test",
		Statement:    "test",
		Answer:       "test",
		Author:       "test",
		Reference:    "test",
		ReferenceURL: "test",
	},
}

var InsertProblem_NoData = []models.Problem{
	{
		// Title:        "test",
		Statement:    "test",
		Answer:       "test",
		Author:       "test",
		Reference:    "test",
		ReferenceURL: "test",
	},
	{
		Title: "test",
		// Statement:    "test",
		Answer:       "test",
		Author:       "test",
		Reference:    "test",
		ReferenceURL: "test",
	},
	{
		Title:     "test",
		Statement: "test",
		// Answer:       "test",
		Author:       "test",
		Reference:    "test",
		ReferenceURL: "test",
	},
	{
		Title:     "test",
		Statement: "test",
		Answer:    "test",
		// Author:       "test",
		Reference:    "test",
		ReferenceURL: "test",
	},
	{
		Title:     "test",
		Statement: "test",
		Answer:    "test",
		Author:    "test",
		// Reference:    "test",
		ReferenceURL: "test",
	},
	{
		Title:     "test",
		Statement: "test",
		Answer:    "test",
		Author:    "test",
		Reference: "test",
		// ReferenceURL: "test",
	},
}

var InsertProblem_Char = []models.Problem{
	{
		Title:        "テスト",
		Statement:    "test",
		Answer:       "test",
		Author:       "test",
		Reference:    "test",
		ReferenceURL: "test",
	},
	{
		Title:        "てすと",
		Statement:    "test",
		Answer:       "test",
		Author:       "test",
		Reference:    "test",
		ReferenceURL: "test",
	},
	{
		Title:        "点検",
		Statement:    "test",
		Answer:       "test",
		Author:       "test",
		Reference:    "test",
		ReferenceURL: "test",
	},
	{
		Title:        "ｱｲｳｴｵ",
		Statement:    "test",
		Answer:       "test",
		Author:       "test",
		Reference:    "test",
		ReferenceURL: "test",
	},
	{
		Title:        `!"#$%&'()=-~^：:；＠」「・。、￥＾ー|\[{]}]*/?.>,<\_`,
		Statement:    "`",
		Answer:       "test",
		Author:       "test",
		Reference:    "test",
		ReferenceURL: "test",
	},
	{
		Title:        "　　　　　    ",
		Statement:    "test",
		Answer:       "test",
		Author:       "test",
		Reference:    "test",
		ReferenceURL: "test",
	},
	{
		Title:        "✲𝑩𝑰𝑮 𝑳𝑶𝑽𝑬😎𖧡",
		Statement:    "test",
		Answer:       "test",
		Author:       "test",
		Reference:    "test",
		ReferenceURL: "test",
	},
}

var InsertProblem_TooLong = []models.Problem{
	{
		// 666 char count
		Title:        "tttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttt",
		Statement:    "test",
		Answer:       "test",
		Author:       "test",
		Reference:    "test",
		ReferenceURL: "test",
	},
	{
		Title:        "test",
		Statement:    "test",
		Answer:       "test",
		Author:       "tttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttt",
		Reference:    "test",
		ReferenceURL: "test",
	},
	{
		Title:        "test",
		Statement:    "test",
		Answer:       "test",
		Author:       "test",
		Reference:    "tttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttt",
		ReferenceURL: "test",
	},
	{
		Title:        "test",
		Statement:    "test",
		Answer:       "test",
		Author:       "test",
		Reference:    "test",
		ReferenceURL: "tttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttt",
	},
}

var InsertProblem_Long = []models.Problem{
	{
		Title:        "test",
		Statement:    "tttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttt",
		Answer:       "test",
		Author:       "test",
		Reference:    "test",
		ReferenceURL: "test",
	},
	{
		Title:        "test",
		Statement:    "test",
		Answer:       "tttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttt",
		Author:       "test",
		Reference:    "test",
		ReferenceURL: "test",
	},
}
