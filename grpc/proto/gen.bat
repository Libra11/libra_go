@echo off
:menu
cls
echo.
echo 1. Run user.proto
echo 2. Run admin.proto
echo 3. Run blog.proto
echo 4. Exit
echo.
set /p choice="Enter your choice: "
if "%choice%"=="1" goto user
if "%choice%"=="2" goto admin
if "%choice%"=="3" goto blog
if "%choice%"=="4" goto end
goto menu

:user
protoc --go_out=../service/user --go_opt=paths=source_relative --go-grpc_out=../service/user --go-grpc_opt=paths=source_relative  ./user.proto
goto menu

:admin
protoc --go_out=../service/admin --go_opt=paths=source_relative --go-grpc_out=../service/admin --go-grpc_opt=paths=source_relative  ./admin.proto
goto menu

:blog
protoc --go_out=../service/blog --go_opt=paths=source_relative --go-grpc_out=../service/blog --go-grpc_opt=paths=source_relative  ./blog.proto
goto menu

:end
exit
