import constants from "./constants"

const initialState = {
    showSidebar: window.innerWidth > 1282,
    showProfilePopup: false,
    showNotificationPopup: false,
    serviceWorkerReady: false
};

export default function reducer(state = initialState, action){
    switch(action.type){
        case constants.TOGGLE_PROFILE_POPUP:
            return Object.assign({}, state, {showProfilePopup: action.payload});
        case constants.TOGGLE_NOTIFICATION_POPUP:
            return Object.assign({}, state, {showNotificationPopup: action.payload});
        case constants.TOGGLE_SIDEBAR:
            return Object.assign({}, state, {showSidebar: action.payload});
        case constants.SET_SERVICE_WORKER_READY:
            return Object.assign({}, state, {serviceWorkerReady: true});
        default:
            return state;
    }
}