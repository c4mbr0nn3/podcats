import equals from 'validator/es/lib/equals'
import isEmail from 'validator/es/lib/isEmail'
import isEmpty from 'validator/es/lib/isEmpty'

export function isEmailRule(value) {
  return isEmail(value) || 'Invalid email.'
}

export function isRequiredRule(value) {
  return !isEmpty(value) || 'Required.'
}

export function isAlphaWithSpacesRule(value) {
  const regex = /^[a-z ]*$/i
  return regex.test(value) || 'Only letters are allowed.'
}

export function confirmPasswordRule(value, password) {
  return equals(value, password) || 'Passwords do not match.'
}

export function validPasswordRule(value) {
  const regex = /^(?=.*[A-Z])(?=.*\d)[A-Z\d]{8,}$/i
  return (
    regex.test(value)
    || 'Password must be at least 8 characters with letters and numbers.'
  )
}
