// This file is generated by "./lib/proto/generate"

package proto

/*

FedCm

This domain allows interacting with the FedCM dialog.

*/

// FedCmLoginState Whether this is a sign-up or sign-in action for this account, i.e.
// whether this account has ever been used to sign in to this RP before.
type FedCmLoginState string

const (
	// FedCmLoginStateSignIn enum const.
	FedCmLoginStateSignIn FedCmLoginState = "SignIn"

	// FedCmLoginStateSignUp enum const.
	FedCmLoginStateSignUp FedCmLoginState = "SignUp"
)

// FedCmDialogType The types of FedCM dialogs.
type FedCmDialogType string

const (
	// FedCmDialogTypeAccountChooser enum const.
	FedCmDialogTypeAccountChooser FedCmDialogType = "AccountChooser"

	// FedCmDialogTypeAutoReauthn enum const.
	FedCmDialogTypeAutoReauthn FedCmDialogType = "AutoReauthn"

	// FedCmDialogTypeConfirmIdpLogin enum const.
	FedCmDialogTypeConfirmIdpLogin FedCmDialogType = "ConfirmIdpLogin"

	// FedCmDialogTypeError enum const.
	FedCmDialogTypeError FedCmDialogType = "Error"
)

// FedCmDialogButton The buttons on the FedCM dialog.
type FedCmDialogButton string

const (
	// FedCmDialogButtonConfirmIdpLoginContinue enum const.
	FedCmDialogButtonConfirmIdpLoginContinue FedCmDialogButton = "ConfirmIdpLoginContinue"

	// FedCmDialogButtonErrorGotIt enum const.
	FedCmDialogButtonErrorGotIt FedCmDialogButton = "ErrorGotIt"

	// FedCmDialogButtonErrorMoreDetails enum const.
	FedCmDialogButtonErrorMoreDetails FedCmDialogButton = "ErrorMoreDetails"
)

// FedCmAccountURLType The URLs that each account has.
type FedCmAccountURLType string

const (
	// FedCmAccountURLTypeTermsOfService enum const.
	FedCmAccountURLTypeTermsOfService FedCmAccountURLType = "TermsOfService"

	// FedCmAccountURLTypePrivacyPolicy enum const.
	FedCmAccountURLTypePrivacyPolicy FedCmAccountURLType = "PrivacyPolicy"
)

// FedCmAccount Corresponds to IdentityRequestAccount.
type FedCmAccount struct {
	// AccountID ...
	AccountID string `json:"accountId"`

	// Email ...
	Email string `json:"email"`

	// Name ...
	Name string `json:"name"`

	// GivenName ...
	GivenName string `json:"givenName"`

	// PictureURL ...
	PictureURL string `json:"pictureUrl"`

	// IdpConfigURL ...
	IdpConfigURL string `json:"idpConfigUrl"`

	// IdpLoginURL ...
	IdpLoginURL string `json:"idpLoginUrl"`

	// LoginState ...
	LoginState FedCmLoginState `json:"loginState"`

	// TermsOfServiceURL (optional) These two are only set if the loginState is signUp
	TermsOfServiceURL string `json:"termsOfServiceUrl,omitempty"`

	// PrivacyPolicyURL (optional) ...
	PrivacyPolicyURL string `json:"privacyPolicyUrl,omitempty"`
}

// FedCmEnable ...
type FedCmEnable struct {
	// DisableRejectionDelay (optional) Allows callers to disable the promise rejection delay that would
	// normally happen, if this is unimportant to what's being tested.
	// (step 4 of https://fedidcg.github.io/FedCM/#browser-api-rp-sign-in)
	DisableRejectionDelay *bool `json:"disableRejectionDelay,omitempty"`
}

// ProtoReq name.
func (m FedCmEnable) ProtoReq() string { return "FedCm.enable" }

// Call sends the request.
func (m FedCmEnable) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// FedCmDisable ...
type FedCmDisable struct{}

// ProtoReq name.
func (m FedCmDisable) ProtoReq() string { return "FedCm.disable" }

// Call sends the request.
func (m FedCmDisable) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// FedCmSelectAccount ...
type FedCmSelectAccount struct {
	// DialogID ...
	DialogID string `json:"dialogId"`

	// AccountIndex ...
	AccountIndex int `json:"accountIndex"`
}

// ProtoReq name.
func (m FedCmSelectAccount) ProtoReq() string { return "FedCm.selectAccount" }

// Call sends the request.
func (m FedCmSelectAccount) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// FedCmClickDialogButton ...
type FedCmClickDialogButton struct {
	// DialogID ...
	DialogID string `json:"dialogId"`

	// DialogButton ...
	DialogButton FedCmDialogButton `json:"dialogButton"`
}

// ProtoReq name.
func (m FedCmClickDialogButton) ProtoReq() string { return "FedCm.clickDialogButton" }

// Call sends the request.
func (m FedCmClickDialogButton) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// FedCmOpenURL ...
type FedCmOpenURL struct {
	// DialogID ...
	DialogID string `json:"dialogId"`

	// AccountIndex ...
	AccountIndex int `json:"accountIndex"`

	// AccountURLType ...
	AccountURLType FedCmAccountURLType `json:"accountUrlType"`
}

// ProtoReq name.
func (m FedCmOpenURL) ProtoReq() string { return "FedCm.openUrl" }

// Call sends the request.
func (m FedCmOpenURL) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// FedCmDismissDialog ...
type FedCmDismissDialog struct {
	// DialogID ...
	DialogID string `json:"dialogId"`

	// TriggerCooldown (optional) ...
	TriggerCooldown *bool `json:"triggerCooldown,omitempty"`
}

// ProtoReq name.
func (m FedCmDismissDialog) ProtoReq() string { return "FedCm.dismissDialog" }

// Call sends the request.
func (m FedCmDismissDialog) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// FedCmResetCooldown Resets the cooldown time, if any, to allow the next FedCM call to show
// a dialog even if one was recently dismissed by the user.
type FedCmResetCooldown struct{}

// ProtoReq name.
func (m FedCmResetCooldown) ProtoReq() string { return "FedCm.resetCooldown" }

// Call sends the request.
func (m FedCmResetCooldown) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// FedCmDialogShown ...
type FedCmDialogShown struct {
	// DialogID ...
	DialogID string `json:"dialogId"`

	// DialogType ...
	DialogType FedCmDialogType `json:"dialogType"`

	// Accounts ...
	Accounts []*FedCmAccount `json:"accounts"`

	// Title These exist primarily so that the caller can verify the
	// RP context was used appropriately.
	Title string `json:"title"`

	// Subtitle (optional) ...
	Subtitle string `json:"subtitle,omitempty"`
}

// ProtoEvent name.
func (evt FedCmDialogShown) ProtoEvent() string {
	return "FedCm.dialogShown"
}

// FedCmDialogClosed Triggered when a dialog is closed, either by user action, JS abort,
// or a command below.
type FedCmDialogClosed struct {
	// DialogID ...
	DialogID string `json:"dialogId"`
}

// ProtoEvent name.
func (evt FedCmDialogClosed) ProtoEvent() string {
	return "FedCm.dialogClosed"
}
