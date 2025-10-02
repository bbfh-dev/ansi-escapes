package escapes

// https://www.ascii-code.com/
const (
	// ASCII control characters (character code 0-31)
	AsciiNull = iota
	AsciiStartOfHeading
	AsciiStartOfText
	AsciiEndOfText
	AsciiEndOfTransmission
	AsciiEnquiry
	AsciiAcknowledge
	AsciiBell
	AsciiBackspace
	AsciiHorizontalTab
	AsciiLineFeed
	AsciiVerticalTab
	AsciiFormFeed
	AsciiCarriageReturn
	AsciiShiftOut
	AsciiShiftIn
	AsciiDataLinkEscape
	AsciiDeviceControlOne
	AsciiDeviceControlTwo
	AsciiDeviceControlThree
	AsciiDeviceControlFour
	AsciiNegativeAcknowledge
	AsciiSyncIdle
	AsciiEndOfTransmissionBlock
	AsciiCancel
	AsciiEndOfMedium
	AsciiSubstitute
	AsciiEscape
	AsciiFileSeparator
	AsciiGroupSeparator
	AsciiRecordSeparator
	AsciiUnitSeparaTOR

	AsciiDelete = 127
)
