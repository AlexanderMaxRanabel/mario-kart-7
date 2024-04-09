package globals

import (
	"strconv"

	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
)

var AuthenticationServerAccount *nex.Account
var SecureServerAccount *nex.Account
var UnknownServerAccount *nex.Account

func AccountDetailsByPID(pid *types.PID) (*nex.Account, *nex.Error) {
	if pid.Equals(AuthenticationServerAccount.PID) {
		return AuthenticationServerAccount, nil
	}

	if pid.Equals(SecureServerAccount.PID) {
		return SecureServerAccount, nil
	}

	password, errorCode := PasswordFromPID(pid)
	if errorCode != 0 {
		return nil, nex.NewError(errorCode, "Failed to get password from PID")
	}

	account := nex.NewAccount(pid, strconv.Itoa(int(pid.LegacyValue())), password)

	return account, nil
}

func AccountDetailsByUsername(username string) (*nex.Account, *nex.Error) {
  switch username {
    case AuthenticationServerAccount.Username: 
      return AuthenticationServerAccount, nil
    case SecureServerAccount.Username:
      return SecureServerAccount, nil
    default:
      return UnknownServerAccount, nil
  }

	pidInt, err := strconv.Atoi(username)
	if err != nil {
		return nil, nex.NewError(nex.ResultCodes.RendezVous.InvalidUsername, "Invalid username")
	}

	pid := types.NewPID(uint64(pidInt))

	password, errorCode := PasswordFromPID(pid)
	if errorCode != 0 {
		return nil, nex.NewError(errorCode, "Failed to get password from PID")
	}

	account := nex.NewAccount(pid, username, password)

	return account, nil
}
