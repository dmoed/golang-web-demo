import constants from './constants'

export function toggleProfilePopup(bool){
    return {type: constants.TOGGLE_PROFILE_POPUP, payload: bool}
}

export function toggleNotificationPopup(bool){
    return {type: constants.TOGGLE_NOTIFICATION_POPUP, payload: bool}
}

export function toggleSidebar(bool){
    return {type: constants.TOGGLE_SIDEBAR, payload: bool}
}

export function updateUser(user){
    return {type: constants.UPDATE_USER, payload: user}
}

export function setServiceWorkerReady(){
    return {type: constants.SET_SERVICE_WORKER_READY}
}