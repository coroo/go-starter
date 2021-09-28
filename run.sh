main_question(){
    echo "============================================\n"
    echo "Please tell me what you want:"
    echo "1. Set Up Environment (.env)"
    echo "2. Do Unit Test"
    echo "3. Update Swagger Documentation"
    echo "4. Run Go Starter\n"
    echo "0. Exit\n\n"
    echo "Your Answer:"
}

bye_statement(){
    echo 
    echo "See you again!"
}

echo "============================================"
echo "Hi, `whoami`\n. Welcome to Go Starter Configuration"
main_question
while :
do
  read INPUT_STRING
  case $INPUT_STRING in
	1)
        echo
        FILE=.env
        if [ ! -f "$FILE" ]; then
            cp .env.example .env
            echo "============================================"
            echo "= Environment file (.env) has been created ="
        fi
        echo "============================================"
        echo "Open Environment in: "
        echo "1. Default IDE"
        echo "2. Visual Code"
        echo "3. Folder\n"
        echo "0. Back\n"
        echo "Your Answer:"
        while :
        do
            read IDE
            case $IDE in
                1)
                    echo "\nOpen Environment in Default IDE ..."
                    open .env
                    bye_statement
                    exit
                    ;;
                2)
                    echo "\nOpen Environment in Visual Code ..."
                    code .env
                    bye_statement
                    exit
                    ;;
                3)
                    echo "\nOpen Environment in Folder ..."
                    open .
                    bye_statement
                    exit
                    ;;
                0)
                    echo
                    main_question
                    break
                    ;;
            esac
        done
        ;;
	2)
        echo "Please wait ..."
        echo "============================================"
        echo "Run for: "
        echo "1. Default Unit Test"
        echo "2. Coverage Unit Test"
        echo "0. Back\n"
        echo "Your Answer:"
        while :
        do
            read IDE
            case $IDE in
                1)
                    echo "Default Unit Test"
                    go test ./... -p=1
                    rm -rf test.db
                    rm -rf testsy.db
                    bye_statement
                    exit
                    ;;
                2)
                    echo "Coverage Unit Test"
                    go test -p=1 -coverprofile cp.out -v ./... && go tool cover -html=cp.out
                    rm -rf test.db
                    rm -rf testsy.db
                    bye_statement
                    exit
                    ;;
                0)
                    echo
                    main_question
                    break
                    ;;
            esac
        done
		;;
	3)
        { # try 1 // Running Swag
            echo
            swag init --parseDependency --parseInternal
        } || { # catch 1
            { # try 2 // Running Swag
                echo "============================================"
                echo "swag not found in app directory ..."
                echo "trying to find swag in GOPATH root folder ..."
                $HOME/go/bin/swag init --parseDependency --parseInternal
            } || { # catch 2 and else
                echo "============================================"
                echo "Please provide your golang location:"
                echo "(e.g. /Users/kuncoro.barot/go)"
                read MY_NAME
                echo "Please wait ..."
                $MY_NAME/bin/swag init --parseDependency --parseInternal
            }
        }
        break
		;;
	4)
        echo "Please wait ..."
		go run main.go
        break
		;;
	0)
		break
		;;
	*)
        echo "============================================"
		echo "Sorry, I don't understand"
		echo "Please answer with correct key or type 0 for close the configuration"
        main_question
		;;
  esac
done
bye_statement