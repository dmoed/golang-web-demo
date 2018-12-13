/**
 * Created by wardo on 3/14/2018.
 */
import React from 'react'
import {Link} from 'react-router-dom'
import onClickOutside from 'react-onclickoutside'
import './style.scss'

class ProfilePopup extends React.Component {

    render() {

        const {profile_image, displayname, email, profile_path, onLogout, exit_impersonation = false} = this.props;

        return (
            <div className="profile-popup">
                <div className="profile-box-head">
                    <div className="profile-box-head-left">
                        <a href={ profile_path }>{profile_image
                            ? <img src={profile_image} className="image"
                                   style={{width: "80px", height: '80px', borderRadius: '80px'}}/>
                            : <span
                            className="preview">{displayname && displayname.length > 0 ? displayname[0] : '...'}</span>}

                        </a>
                    </div>
                    <div className="profile-box-head-right">
                        <span className="bold">{displayname}</span>
                        <span className="small">{email}</span>
                    </div>
                </div>
                <div className="profile-box-bottom">
                    <ul>
                        {/*<li>
                            <Link to="/account/change-password"><span className="mdi mdi-security"></span>Change Password</Link>
                        </li>*/}
                        <li>
                            <a href="#logout" onClick={onLogout}><span className="mdi mdi-logout"></span>Sign out</a>
                        </li>
                    </ul>
                </div>
            </div>
        );
    }

    handleClickOutside() {
        if (this.props.close) {
            this.props.close();
        }
    }
}

export default onClickOutside(ProfilePopup)
