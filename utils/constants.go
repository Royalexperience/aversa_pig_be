package utilsConstants

const QUERY_ALL_DATA = "SELECT email, username, password FROM utenti"
const QUERY_INSERT_NEW_DATA = "INSERT INTO utenti (username, email, password) VALUES ($1, $2, $3)"
const EMAIL_FIELD_MISSING = "Email field is missing"
const USERNAME_FIELD_MISSING = "Username field is missing"
const PASSWORD_FIELD_MISSING = "Password field is missing"
const ERROR_PARSING_RESPONSE= "Error Parsing Response"
const ERROR_400_BAD_REQUEST= ERROR_400+"Bad Request"
const ERROR_405_METHOD_NOT_ALLOWED= ERROR_405+" "+ METHOD_NOT_ALLOWED
const ERROR_500_QUERY= ERROR_500+"Error trying to execute query"
const ERROR_400 = "[ERROR 400]"
const ERROR_405 = "[ERROR 405]"
const ERROR_500 = "[ERROR 500]"
const METHOD_NOT_ALLOWED = "Method not allowed"
const METHOD_GET = "GET"
const METHOD_POST = "POST"
const ERROR_MESSAGE_EMPTY_BODY_OR_WRONG_FIELDS = "Error request results as an empty Body or with wrong field names"