package utilsConstants

const QUERY_ALL_DATA = "SELECT email, username, password FROM utenti"
const QUERY_INSERT_NEW_DATA = "INSERT INTO utenti (username, email, password) VALUES ($1, $2, $3)"