/**
 * Created by wardo on 3/14/2018.
 */

import React from 'react'
import './style.scss'

const Avatar = ({displayname, image, onClick}) => {
    return (
        <div className="avatar" onClick={onClick}>
            <a className="avatar-rnd nav-link">
                {displayname && displayname.length > 0 ? displayname[0] : ""}
            </a>
            <span className="avatar-displayname">{displayname}</span>
        </div>
    );
};

export default Avatar