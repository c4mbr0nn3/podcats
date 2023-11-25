import isEmail from 'validator/es/lib/isEmail'
import isEmpty from 'validator/es/lib/isEmpty'
import equals from 'validator/es/lib/equals'

export function isEmailRule(value) {
  return isEmail(value) || 'Invalid email.'
}

export function isRequiredRule(value) {
  return !isEmpty(value) || 'Required.'
}

export function isAlphaWithSpacesRule(value) {
  const regex = /^[a-zA-Z ]*$/
  return regex.test(value) || 'Only letters are allowed.'
}

export function confirmPasswordRule(value, password) {
  return equals(value, password) || 'Passwords do not match.'
}

export function validPasswordRule(value) {
  const regex = /^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{8,}$/
  return (
    regex.test(value)
    || 'Password must be at least 8 characters with letters and numbers.'
  )
}
