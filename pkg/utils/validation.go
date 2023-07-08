package utils

func ValidationIsExpired(createdTimestamp int, expiresTimeInMinutes int) bool {
	expiresTimeInMS := ConvertMinutesToMS(expiresTimeInMinutes)

	now := TimeNowInTimestamp()
	validationExpiresTimestamp := createdTimestamp + expiresTimeInMS

	return now >= validationExpiresTimestamp
}
