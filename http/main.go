package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	server := &http.Server{
		Addr: ":9091",
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("Request:", r.URL.String())
		_, err := w.Write([]byte(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>forum</title>
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@4.5.0/dist/css/bootstrap.min.css"
          integrity="sha384-9aIt2nRpC12Uk9gS9baDl411NQApFmC26EwAOH8WgZl5MYYxFfc+NcPb1dKGj7Sk" crossorigin="anonymous">
    <script src="https://s3.pstatp.com/cdn/expire-1-M/jquery/3.3.1/jquery.min.js"></script>
    <script src="https://cdn.jsdelivr.net/npm/popper.js@1.16.0/dist/umd/popper.min.js"
            integrity="sha384-Q6E9RHvbIyZFJoft+2mJbHaEWldlvI9IOYy5n3zV9zzTtmI3UksdQRVvoxMfooAo"
            crossorigin="anonymous"></script>
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@4.5.0/dist/js/bootstrap.min.js"
            integrity="sha384-OgVRvuATP1z7JjHLkuOU7Xw704+h835Lr+6QL9UvYjZE3Ipu6Tp75j7Bh/kR0JKI"
            crossorigin="anonymous"></script>
    <style>
        .titleFont {
            font-weight: bold;
            font-family: Arial, serif;
            font-style: italic;
        }

        .logFont {
            margin-top: 5px;
            font-weight: bold;
            font-size: 13px;
            font-family: "Arial Black", serif;
        }
    </style>
    <script src="js/userinfo/logReg.js"></script>

</head>
<body>
<nav class="navbar navbar-expand-lg navbar-light bg-light">
    <a class="navbar-brand titleFont" href="./inside_page.html" style="color: red">Sunshine</a>

    <div class="collapse navbar-collapse" id="navbarSupportedContent">
        <ul class="navbar-nav mr-auto">
            <li class="nav-item active">
                <a class="nav-link" href="./inside_page.html">Home <span class="sr-only">(current)</span></a>
            </li>
            <li class="nav-item">
                <a class="nav-link" href="#">Link</a>
            </li>
        </ul>
        <form class="form-inline my-2 my-lg-0">
            <input class="form-control mr-sm-2" type="search" placeholder="Search" aria-label="Search">
            <button class="btn btn-outline-success my-2 my-sm-0" type="submit" id="searchTopic">Search</button>
        </form>
    </div>
</nav>
<div class="titleFont">
    hahaha
</div>
<img src="https://th.bing.com/th/id/R4e5ba6c7d897542e2ba8f6c848b65a45?rik=76Hxw3hHHK8xpQ&riu=http%3a%2f%2fimg.xidong.net%2fxdpic_n%2f7011511%2f1156905.jpg&ehk=D3aWkM32b%2bQQq6I%2fZ07Vf2%2fL60mggl5stuFvS%2bZuDcQ%3d&risl=&pid=ImgRaw">
<div class="container-xl">
    <div class="row">
        <div class="col-8 " style="min-height: 600px">
            <div id="createTopicDiv" style="display: none">
                <form class="p-2" id="topicFrom">
                    <div class="form-group">
                        <label for="topicInput">Create New Topic ! </label>
                        <input type="text" class="form-control" id="topicInput">
                    </div>
                    <div class="d-flex justify-content-around" id="category_list">
                        <div class="form-check form-check-inline">
                            <input class="form-check-input" type="radio" name="inlineRadioOptions" id="inlineRadio1"
                                   value="option1">
                            <label class="form-check-label" for="inlineRadio1">life</label>
                        </div>
                        <div class="form-check form-check-inline">
                            <input class="form-check-input" type="radio" name="inlineRadioOptions" id="inlineRadio2"
                                   value="option2">
                            <label class="form-check-label" for="inlineRadio2">work</label>
                        </div>
                        <div class="form-check form-check-inline">
                            <input class="form-check-input" type="radio" name="inlineRadioOptions" id="inlineRadio3"
                                   value="option3">
                            <label class="form-check-label" for="inlineRadio3"> study</label>
                        </div>
                        <div class="form-check form-check-inline">
                            <input class="form-check-input" type="radio" name="inlineRadioOptions" id="inlineRadio4"
                                   value="option3">
                            <label class="form-check-label" for="inlineRadio3"> family</label>
                        </div>
                        <div class="form-check form-check-inline">
                            <input class="form-check-input" type="radio" name="inlineRadioOptions" id="inlineRadio5"
                                   value="option3">
                            <label class="form-check-label" for="inlineRadio3"> else</label>
                        </div>
                        <div class="btn btn-primary " id="createBtn">create</div>
                    </div>
                </form>
            </div>
            <div id="topicListDiv">

            </div>
        </div>
        <div class="col-4">
            <div class="d-flex flex-column">
                <form class="p-2" id="login_form">
                    <div class="form-group">
                        <label for="username">Username</label>
                        <input type="text" class="form-control" id="username">
                    </div>
                    <div class="form-group">
                        <label for="password">Password</label>
                        <input type="password" class="form-control" id="password">
                    </div>
                    <div class="btn btn-primary" id="loginBtn">Login</div>
                    <div class="d-flex justify-content-between">
                        <p class="logFont">forget password</p>
                        <p class="logFont" id="registerBtn">register</p>
                    </div>
                </form>
                <div style="display:none;" id="userInfoVo">
                    <div class="p-2 d-flex flex-column justify-content-center"
                         style="min-height: 300px;">
                        <div style="text-align: center"><img class="user_avator" src=""/></div>
                        <div id="usernameDiv" style="text-align: center;margin-top: 10px">

                        </div>
                    </div>
                </div>
<!--                <div class="p-2 bg-secondary" style="min-height: 300px">hot topic</div>-->
            </div>
        </div>
    </div>
</div>


</body>
<script src="js/userinfo/common.js"></script>
<script src="js/userinfo/logReg.js"></script>
<script>
    $(function () {
        $.getUserInfo(function (userInfo) {
            console.log($isLogin)
            if ($isLogin) {
                $("#userInfoVo .user_avator").attr("src",userInfo.avatar)
                $("#usernameDiv").html("<h4>"+userInfo.username+"</h4>")
                $("#login_form").hide()
                $("#userInfoVo").show()
                $("#createTopicDiv").show()
                console.log(userInfo)
            } else {
                $("#login_form").show()
                $("#userInfoVo").hide()
                $("#createTopicDiv").hide()
            }
        })

    })
</script>
<script src="js/category/category.js"></script>

</html>`))
		if err != nil {
			fmt.Println(err)
		}
	})

	server.Handler = handler

	log.Print("starting HTTP server on port: 3000")

	server.ListenAndServe()

	ch := make(chan int)
	<-ch
}
