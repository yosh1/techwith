
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/jinzhu/gorm"
	// for debuging
	"fmt"
	"reflect"
	"os/exec"
	"os"
	// for casting
	"unsafe"
	"ps-editor/common"
	"ps-editor/controllers"
	"ps-editor/models"
	"strings"

)

func runProgram(data string) string {
	file, err := os.Create("program.php")
	if err != nil {
		fmt.Println(err)
	}
	file.Write(([]byte)(data))
	output, err := exec.Command("php", "program.php").CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	return *(*string)(unsafe.Pointer(&output))
}

func migrate(db *gorm.DB) {

	// create scheme
	db.AutoMigrate(models.User{}, models.Question{}, models.Answer{}, models.Level{})

	// unit 1
	db.Create(models.Question{
		Id: 1,
		Title: "変数",
		Body: "siteという変数に\"kirakira\"という文字列を代入し、値をコンソールに出力して下さい",
		Hint: "変数の先頭には$マークをつけるのを忘れないで下さい",
		LevelId: 2,
	})
	db.Create(models.Question{
		Id: 2,
		Title: "変数",
		Body: "siteという変数を宣言して下さい",
		Hint: "変数の宣言$変数名とします",
		LevelId: 1,
	})
	db.Create(models.Question{
		Id: 3,
		Title: "変数",
		Body: "コンソールに下記のように出力して下さい。ただし、kirakiraは変数に代入して表示させて下さい。hello kirakira",
		Hint: "変数の先頭には$マークをつけるのを忘れないで下さい",
		LevelId: 3,
	})

	// // unit 2
	// db.Create(models.Question{
	// 	Title: "定数",
	// 	Body: "siteという変数に\"kirakira\"という文字列を代入し、値をコンソールに出力して下さい。",
	// 	Hint: "define('SITE', 'kirakira')",
	// 	LevelId: 2,
	// })
	// db.Create(models.Question{
	// 	Title: "定数",
	// 	Body: "siteという変数に\"kirakira\"という文字列を代入し、値をコンソールに出力して下さい。",
	// 	Hint: "define('SITE', 'kirakira')",
	// 	LevelId: 1,
	// })
	// db.Create(models.Question{
	// 	Title: "定数",
	// 	Body: "siteという変数に\"kirakira\"という文字列を代入し、値をコンソールに出力して下さい。",
	// 	Hint: "define('SITE', 'kirakira')",
	// 	LevelId: 3,
	// })
	//
	// // unit 3
	// db.Create(models.Question{
	// 	Id: 7, Title: "配列",
	// 	Body: "fruitという配列を宣言し、それに'banana', 'apple', 'orange'を代入して下さい",
	// 	Hint: "配列の作成にはarray()を使用します。",
	// 	LevelId: 2,
	// })
	// db.Create(models.Question{
	// 	Id: 8, Title: "配列",
	// 	Body: "siteという変数に\"kirakira\"という文字列を代入し、値をコンソールに出力して下さい。",
	// 	Hint: "配列の作成にはarray()を使用します。",
	// 	LevelId: 1,
	// })
	// db.Create(models.Question{
	// 	Id: 9, Title: "配列",
	// 	Body: "siteという変数に\"kirakira\"という文字列を代入し、値をコンソールに出力して下さい。",
	// 	Hint: "配列の作成にはarray()を使用します。",
	// 	LevelId: 3,
	// })
	//
	// // unit 4 連想配列
	// db.Create(models.Question{
	// 	Id: 11, Title: "連想配列",
	// 	Body: "fruitという配列を宣言し、それに'banana', 'apple', 'orange'を代入して下さい",
	// 	Hint: "配列の作成にはarray()を使用します。",
	// 	LevelId: 2,
	// })
	// db.Create(models.Question{
	// 	Id: 12, Title: "連想配列",
	// 	Body: "siteという変数に\"kirakira\"という文字列を代入し、値をコンソールに出力して下さい。",
	// 	Hint: "変数の先頭には$マークをつけるのを忘れないで下さい",
	// 	LevelId: 1,
	// })
	// db.Create(models.Question{
	// 	Id: 13, Title: "連想配列",
	// 	Body: "siteという変数に\"kirakira\"という文字列を代入し、値をコンソールに出力して下さい。",
	// 	Hint: "変数の先頭には$マークをつけるのを忘れないで下さい",
	// 	LevelId: 3,
	// })
	//
	// // unit 5 2次元配列
	// db.Create(models.Question{
	// 	Id: 14, Title: "2次元配列",
	// 	Body: "fruitという配列を宣言し、それに'banana', 'apple', 'orange'を代入して下さい",
	// 	Hint: "配列の作成にはarray()を使用します。",
	// 	LevelId: 2,
	// })
	// db.Create(models.Question{
	// 	Id: 15, Title: "2次元配列",
	// 	Body: "siteという変数に\"kirakira\"という文字列を代入し、値をコンソールに出力して下さい。",
	// 	Hint: "変数の先頭には$マークをつけるのを忘れないで下さい",
	// 	LevelId: 1,
	// })
	// db.Create(models.Question{
	// 	Id: 16, Title: "2次元配列",
	// 	Body: "siteという変数に\"kirakira\"という文字列を代入し、値をコンソールに出力して下さい。",
	// 	Hint: "変数の先頭には$マークをつけるのを忘れないで下さい",
	// 	LevelId: 3,
	// })
	//
	// // unit 6 演算子
	// db.Create(models.Question{
	// 	Id: 17, Title: "演算子",
	// 	Body: "fruitという配列を宣言し、それに'banana', 'apple', 'orange'を代入して下さい",
	// 	Hint: "配列の作成にはarray()を使用します。",
	// 	LevelId: 2,
	// })
	// db.Create(models.Question{
	// 	Id: 18, Title: "演算子",
	// 	Body: "siteという変数に\"kirakira\"という文字列を代入し、値をコンソールに出力して下さい。",
	// 	Hint: "変数の先頭には$マークをつけるのを忘れないで下さい",
	// 	LevelId: 1,
	// })
	// db.Create(models.Question{
	// 	Id: 19, Title: "演算子",
	// 	Body: "siteという変数に\"kirakira\"という文字列を代入し、値をコンソールに出力して下さい。",
	// 	Hint: "変数の先頭には$マークをつけるのを忘れないで下さい",
	// 	LevelId: 3,
	// })
	//
  // // if文
	// db.Create(models.Question{
	// 	Id: 20, Title: "if文",
	// 	Body: "fruitという配列を宣言し、それに'banana', 'apple', 'orange'を代入して下さい",
	// 	Hint: "配列の作成にはarray()を使用します。",
	// 	LevelId: 2,
	// })
	// db.Create(models.Question{
	// 	Id: 21, Title: "if文",
	// 	Body: "siteという変数に\"kirakira\"という文字列を代入し、値をコンソールに出力して下さい。",
	// 	Hint: "変数の先頭には$マークをつけるのを忘れないで下さい",
	// 	LevelId: 1,
	// })
	// db.Create(models.Question{
	// 	Id: 22, Title: "if文",
	// 	Body: "siteという変数に\"kirakira\"という文字列を代入し、値をコンソールに出力して下さい。",
	// 	Hint: "変数の先頭には$マークをつけるのを忘れないで下さい",
	// 	LevelId: 3,
	// })
	//
	// // switch文
	// db.Create(models.Question{
	// 	Id: 23, Title: "switch文",
	// 	Body: "fruitという配列を宣言し、それに'banana', 'apple', 'orange'を代入して下さい",
	// 	Hint: "配列の作成にはarray()を使用します。",
	// 	LevelId: 2,
	// })
	// db.Create(models.Question{
	// 	Id: 24, Title: "switch文",
	// 	Body: "siteという変数に\"kirakira\"という文字列を代入し、値をコンソールに出力して下さい。",
	// 	Hint: "変数の先頭には$マークをつけるのを忘れないで下さい",
	// 	LevelId: 1,
	// })
	// db.Create(models.Question{
	// 	Id: 25, Title: "switch文",
	// 	Body: "siteという変数に\"kirakira\"という文字列を代入し、値をコンソールに出力して下さい。",
	// 	Hint: "変数の先頭には$マークをつけるのを忘れないで下さい",
	// 	LevelId: 3,
	// })
	//
	// // for文
	// db.Create(models.Question{
	// 	Id: 26, Title: "for文",
	// 	Body: "fruitという配列を宣言し、それに'banana', 'apple', 'orange'を代入して下さい",
	// 	Hint: "配列の作成にはarray()を使用します。",
	// 	LevelId: 2, // middle
	// })
	// db.Create(models.Question{
	// 	Id: 27, Title: "for文",
	// 	Body: "siteという変数に\"kirakira\"という文字列を代入し、値をコンソールに出力して下さい。",
	// 	Hint: "変数の先頭には$マークをつけるのを忘れないで下さい",
	// 	LevelId: 1,
	// })
	// db.Create(models.Question{
	// 	Id: 28, Title: "for文",
	// 	Body: "siteという変数に\"kirakira\"という文字列を代入し、値をコンソールに出力して下さい。",
	// 	Hint: "変数の先頭には$マークをつけるのを忘れないで下さい",
	// 	LevelId: 3,
	// })
	//
	// // while文
	// db.Create(models.Question{
	// 	Id: 29, Title: "while文",
	// 	Body: "fruitという配列を宣言し、それに'banana', 'apple', 'orange'を代入して下さい",
	// 	Hint: "配列の作成にはarray()を使用します。",
	// 	LevelId: 2,
	// })
	// db.Create(models.Question{
	// 	Id: 30, Title: "while文",
	// 	Body: "siteという変数に\"kirakira\"という文字列を代入し、値をコンソールに出力して下さい。",
	// 	Hint: "変数の先頭には$マークをつけるのを忘れないで下さい",
	// 	LevelId: 1,
	// })
	// db.Create(models.Question{
	// 	Id: 31, Title: "while文",
	// 	Body: "siteという変数に\"kirakira\"という文字列を代入し、値をコンソールに出力して下さい。",
	// 	Hint: "変数の先頭には$マークをつけるのを忘れないで下さい",
	// 	LevelId: 3,
	// })
	//
	// // foreach
	// db.Create(models.Question{
	// 	Id: 7, Title: "foreach文",
	// 	Body: "fruitという配列を宣言し、それに'banana', 'apple', 'orange'を代入して下さい",
	// 	Hint: "配列の作成にはarray()を使用します。",
	// 	LevelId: 2,
	// })
	// db.Create(models.Question{
	// 	Id: 8, Title: "foreach文",
	// 	Body: "siteという変数に\"kirakira\"という文字列を代入し、値をコンソールに出力して下さい。",
	// 	Hint: "変数の先頭には$マークをつけるのを忘れないで下さい",
	// 	LevelId: 1,
	// })
	// db.Create(models.Question{
	// 	Id: 9, Title: "foreach文",
	// 	Body: "siteという変数に\"kirakira\"という文字列を代入し、値をコンソールに出力して下さい。",
	// 	Hint: "変数の先頭には$マークをつけるのを忘れないで下さい",
	// 	LevelId: 3,
	// })
	//
	// // func
	// db.Create(models.Question{
	// 	Id: 7, Title: "関数",
	// 	Body: "fruitという配列を宣言し、それに'banana', 'apple', 'orange'を代入して下さい",
	// 	Hint: "配列の作成にはarray()を使用します。",
	// 	LevelId: 2,
	// })
	// db.Create(models.Question{
	// 	Id: 8, Title: "関数",
	// 	Body: "siteという変数に\"kirakira\"という文字列を代入し、値をコンソールに出力して下さい。",
	// 	Hint: "変数の先頭には$マークをつけるのを忘れないで下さい",
	// 	LevelId: 1,
	// })
	// db.Create(models.Question{
	// 	Id: 9, Title: "関数",
	// 	Body: "siteという変数に\"kirakira\"という文字列を代入し、値をコンソールに出力して下さい。",
	// 	Hint: "変数の先頭には$マークをつけるのを忘れないで下さい",
	// 	LevelId: 3,
	// })
	//
	// // string utils
	// db.Create(models.Question{
	// 	Id: 7, Title: "文字列の操作関数",
	// 	Body: "fruitという配列を宣言し、それに'banana', 'apple', 'orange'を代入して下さい",
	// 	Hint: "配列の作成にはarray()を使用します。",
	// 	LevelId: 2,
	// })
	// db.Create(models.Question{
	// 	Id: 8, Title: "文字列の操作関数",
	// 	Body: "siteという変数に\"kirakira\"という文字列を代入し、値をコンソールに出力して下さい。",
	// 	Hint: "変数の先頭には$マークをつけるのを忘れないで下さい",
	// 	LevelId: 1,
	// })
	// db.Create(models.Question{
	// 	Id: 9, Title: "文字列の操作関数",
	// 	Body: "siteという変数に\"kirakira\"という文字列を代入し、値をコンソールに出力して下さい。",
	// 	Hint: "変数の先頭には$マークをつけるのを忘れないで下さい",
	// 	LevelId: 3,
	// })
	//
	// // array utils
	// db.Create(models.Question{
	// 	Id: 7, Title: "配列の操作関数",
	// 	Body: "fruitという配列を宣言し、それに'banana', 'apple', 'orange'を代入して下さい",
	// 	Hint: "配列の作成にはarray()を使用します。",
	// 	LevelId: 2,
	// })
	// db.Create(models.Question{
	// 	Id: 8, Title: "配列の操作関数",
	// 	Body: "siteという変数に\"kirakira\"という文字列を代入し、値をコンソールに出力して下さい。",
	// 	Hint: "変数の先頭には$マークをつけるのを忘れないで下さい",
	// 	LevelId: 1,
	// })
	// db.Create(models.Question{
	// 	Id: 9, Title: "配列の操作関数",
	// 	Body: "siteという変数に\"kirakira\"という文字列を代入し、値をコンソールに出力して下さい。",
	// 	Hint: "変数の先頭には$マークをつけるのを忘れないで下さい",
	// 	LevelId: 3,
	// })



	// // etc
	// db.Create(models.Question{
	// 	Id: 1, Title: "echo関数の使い方",
	// 	Body: "コンソールに\"hoge\"と出力して下さい。",
	// 	Hint: "文字列はダブルクォーテーションで囲います。",
	// })
	// db.Create(models.Question{
	// 	Id: 2, Title: "for文の使い方",
	// 	Body: "for文を用いて、コンソールに1から10までの数字を出力して下さい。",
	// 	Hint: "改行コードは、\nであらわします。",
	// })


	db.Create(models.Answer{
		QuestionId: 1,
		Body:"kirakira",
		KeyWord: "$site",
		Message: "$siteという変数を宣言して下さい",
	})
	db.Create(models.Answer{
		QuestionId: 2, Body:"12345678910",
		KeyWord: "$",
		Message: "変数を宣言して下さい",
	})
	db.Create(models.Answer{
		QuestionId: 3, Body:"hello kirakira",
		KeyWord: "$",
		Message: "変数を宣言して下さい",
	})
	// db.Create(models.Answer{QuestionId: 3, Body:"kirakira"})
	// db.Create(models.Answer{QuestionId: 3, Body:"kirakira"})
	// db.Create(models.Answer{QuestionId: 3, Body:"kirakira"})

	db.Create(models.Level{Id: 1, Level:"easy"})
	db.Create(models.Level{Id: 2, Level:"medium"})
	db.Create(models.Level{Id: 3, Level:"hard"})
}

func main() {
	  db := common.Init()
		migrate(db)
		// must
	  defer db.Close()

		r := gin.Default()
		// auth
		// r.Use(common.AuthReq("token"))
		r.LoadHTMLGlob("templates/*")
		r.Static("/assets", "./assets")

		// editor / endpoint
		r.GET("/", func(c *gin.Context) {
			ctrl := controllers.NewQuestion()
			// harded coded, the time being
			question := ctrl.GetId("1")
			fmt.Println(question)
			c.HTML(http.StatusOK, "index.tmpl", gin.H{
				"question": question,
			})
		})

		// exec program
		r.POST("/", func(c *gin.Context) {

			result := true
			message := ""

			data := c.PostForm("data")
			qNum := c.PostForm("qNum")

			// get the keyWord
			ctrl := controllers.NewAnswer()
			keyWord := ctrl.GetKeyWord(qNum)
			if keyWord == "" {
				fmt.Println("keyWord is nil!!")
			} else {
				fmt.Println("keyWord is %s", keyWord)
			}

			if strings.Index(data, keyWord) < 0 {
				result = false
				message = ctrl.GetMessage(qNum)
				fmt.Println("message is ")
				fmt.Println(message)
			}

			console := runProgram(data)

      // DEBUG:
			fmt.Println("output type is ")
			fmt.Println(reflect.TypeOf(console))

			// get the answer
			ctrl = controllers.NewAnswer()
			answer := ctrl.GetBody(qNum)

			// DEBUG:
			fmt.Println("answer type is ")
			fmt.Println(reflect.TypeOf(answer))

			// message := ctrl.GetMessage(qNum)
			// fmt.Println("message is ....")
			// fmt.Println(message)

			// if strings.Index(data, keyWord) < 0 {
			// 	result = false
			// }

			fmt.Println("answer is %s", answer)

			if answer != console {
			  result = false
			}

			// c.HTML(http.StatusOK, "index.tmpl", gin.H{
			// 	"data": data,
			// 	// "output": strings.Trim(string(output), " "),
			c.JSON(200, gin.H{
				//"output": *(*string)(unsafe.Pointer(&output)),
				"output": console,
				"result": result,
				"message": message,
			})
		})

		// run program
		r.POST("/run", func(c *gin.Context) {
			data := c.PostForm("data")
			console := runProgram(data)
			c.JSON(200, gin.H{
				"output": console,
			})
		})

		/*
		r.GET("/question", func(c *gin.Context) {
			//qNum := c.Param("qNum")
			qNum := c.Query("qNum")
			ctrl := controllers.NewQuestion()
			fmt.Println(qNum)
			fmt.Println(reflect.TypeOf(qNum))
			question := ctrl.GetId(qNum)
			fmt.Println(question)

			// c.JSON(http.StatusOK, gin.H{
			// 	"question": question,
			// })
			c.HTML(http.StatusOK, "index.tmpl", gin.H{
				"question": question,
			})
		})
		*/

		r.GET("/question/:id", func(c *gin.Context) {
			id := c.Param("id")
			ctrl := controllers.NewQuestion()
			question := ctrl.GetId(id)
			// fmt.Println(question.Title)
			c.HTML(http.StatusOK, "index.tmpl", gin.H{
				"question": question,
			})
		})

		// 中断
		r.POST("/suspend", func(c *gin.Context) {
			//uId := c.PostForm("uId")
			//qNum := c.PostForm("qNum")
			//ctrl := controllers.NewUser()

			result := true
			c.JSON(http.StatusOK, gin.H{
				"result": result,
			})
		})

		/*
		dead or alive
		*/
		r.GET("/ping", func(c *gin.Context) {
			// harded coded, the time being
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

		r.Run() // listen and serve on 0.0.0.0:8080
	}
