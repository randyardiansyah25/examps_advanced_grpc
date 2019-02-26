package usecase

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/randyardiansyah25/examps_advanced_grpc/grpc_client/repository"
	"github.com/randyardiansyah25/examps_advanced_grpc/grpc_common/model"
	"github.com/randyardiansyah25/libpkg/util/str"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os"
	"strconv"
	"strings"
	"time"
)

type RequestUsecase struct {
	clientRepo model.UserHandlerClient
	reqContext context.Context
	cancelF context.CancelFunc
	timeOut int64
}

func NewRequestUsecase() *RequestUsecase {
	host := fmt.Sprintf("%s:%s", os.Getenv("HOST_ADDR"), os.Getenv("HOST_PORT"))
	sto := os.Getenv("REQUEST_TIMEOUT")
	to,_ := strconv.ParseInt(sto, 10, 64)

	deadLine := time.Now().Add(time.Duration(to) * time.Second)
	ctx, cF := context.WithDeadline(context.Background(), deadLine)
	return &RequestUsecase{
		clientRepo:repository.NewClientRepo(host),
		reqContext:ctx,
		cancelF:cF,
		timeOut:to,
	}
}

func (_r *RequestUsecase) AddUser(id string, name string, pwd string, gender string) string {
	defer _r.cancelF()
	if name == "" {
		return "invalid name of user"
	}

	if pwd == "" {
		return "invalid user password"
	}

	if id == "" {
		return "invalid user id"
	}

	var gen model.UserGender
	if strings.ToLower(gender) == "l"{
		gen = model.UserGender_Male
	}else if strings.ToLower(gender) == "p"{
		gen = model.UserGender_Female
	}else{
		return "invalid user gender!!"
	}

	user := model.User{
		Id:id,
		Name:name,
		Password:pwd,
		Gender:gen,
	}

	resp, err := _r.clientRepo.AddUser(_r.reqContext, &user)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			switch st.Code() {
			case codes.Unavailable:
				//refused or handshake failed
				return fmt.Sprintf("[CONN ERR] Cannot communicate with server : %s", st.Message())
			case codes.DeadlineExceeded:
				//Timeout reached
				return fmt.Sprintf("[CONN ERR] Request timeout after %d : %s", _r.timeOut, st.Message())
			}
		}else{
			return fmt.Sprintf("[ERROR] %s", err.Error())
		}
	}

	return fmt.Sprintf("[%s] %s", resp.ResponseCode, resp.ResponseMsg)
}

func (_r *RequestUsecase) GetUserById(id string) string{
	defer _r.cancelF()

	uid := model.UserId{}
	uid.Id = id
	user, err := _r.clientRepo.GetUser(_r.reqContext, &uid)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			switch st.Code() {
			case codes.Unavailable:
				//refused or handshake failed
				return fmt.Sprintf("[CONN ERR] Cannot communicate with server : %s", st.Message())
			case codes.DeadlineExceeded:
				//Timeout reached
				return fmt.Sprintf("[CONN ERR] Request timeout after %d : %s", _r.timeOut, st.Message())
			}
		}else{
			return fmt.Sprintf("[ERROR] %s", err.Error())
		}
	}
	var rs = []string{
		"ID       : ", user.Id,"\n",
		"Name     : ", user.Name, "\n",
		"Password : ", user.Password, "\n",
		"Gender   : ", _r.getGenderString(user.Gender), "\n",
	}
	
	return strings.Join(rs, "")
}

func (_r *RequestUsecase) GetUsers() string {
	defer _r.cancelF()
	list, err := _r.clientRepo.ListUser(context.Background(), &empty.Empty{})
	if err != nil {
		if st, ok := status.FromError(err); ok {
			switch st.Code() {
			case codes.Unavailable:
				//refused or handshake failed
				return fmt.Sprintf("[CONN ERR] Cannot communicate with server : %s", st.Message())
			case codes.DeadlineExceeded:
				//Timeout reached
				return fmt.Sprintf("[CONN ERR] Request timeout after %d : %s", _r.timeOut, st.Message())
			}
		}else{
			return fmt.Sprintf("[ERROR] %s", err.Error())
		}
	}

	var table = []string{
		"No  ID        Name                   Password        Gender       \n",
		"==================================================================\n",
	//   123412345678901234567890123456789012312345678901234561234567890123
	//       x         x                      x               x
	}
	for i, user := range list.List {
		row := []string{
			strutils.RightPad(strconv.Itoa(i), 4," "),
			strutils.RightPad(user.Id, 10, " "),
			strutils.RightPad(user.Name, 23, " "),
			strutils.RightPad(user.Password, 16, " "),
			strutils.RightPad(_r.getGenderString(user.Gender), 13, " "),
			"\n",
		}
		table = append(table, strings.Join(row, ""))
	}

	return strings.Join(table, "")
}

func (_r *RequestUsecase) getGenderString(gen model.UserGender) string{
	switch gen {
	case model.UserGender_Male:
		return "LAKI-LAKI"
	case model.UserGender_Female:
		return "PEREMPUAN"
	}
	return "UNDEFINED"
}

