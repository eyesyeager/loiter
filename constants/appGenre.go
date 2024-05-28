package constants

/**
 * 应用类型
 * @auth eyesYeager
 * @date 2024/3/6 15:41
 */

var AppGenre = appGenre{
	Api:    "api",
	Static: "static",
}

type appGenre struct {
	Api    string
	Static string
}
