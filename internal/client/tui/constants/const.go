package constants

const (
	StateMainMenu      = "main_menu"
	StateLoginEmail    = "login_email"
	StateLoginPassword = "login_password"
	StateLogout        = "logout"
	StateOtpRequested  = "otp_requested"
	StateOtpVerified   = "otp_verified"
	StateDataMenu      = "data_menu"
	StateSwitchToAuth  = "switch_to_auth"
	StateSwitchToData  = "switch_to_data"
	StateRefreshNeeded = "refresh_needed"
	StateQuit          = "quit"

	StateAuthorizedMainMenu = "authorized_main_menu"
)

const (
	CmdUp        = "up"
	CmdDown      = "down"
	CmdEnter     = "enter"
	CmdBack      = "backspace"
	CmdForceQuit = "ctrl+c"
)
