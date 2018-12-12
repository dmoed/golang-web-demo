import constants from './constants'

const initialState = {
    id: 0,
    username: null,
    displayname: null,
    email: null,
    profile_image: null,
    roles: [],
    authenticated: true
};

export default function reducer(state = initialState, action){
    switch(action.type){

        case constants.UPDATE_USER:
            return Object.assign({}, state, {user: action.payload});

        default:
            return state;
    }
}