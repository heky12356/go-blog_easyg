import React from 'react';
import { Navigate } from 'react-router-dom';
import Cookies from 'js-cookie';
export default function Verify({children}) {
    var refreshtoken = Cookies.get('refreshToken');
    if (!refreshtoken) {
        return <Navigate to="/login" />;
    }
    return (
        <>
            {children}
        </>
    )
}