import React from 'react';
import { Outlet, useLocation, Link} from 'react-router-dom';
export default function Admin () {
    const location = useLocation(); // 获取当前路径
    let titleSuffix = ""; // 初始化动态后缀为空
    // 动态解析路径后缀，只取整个路径中的第一级
    if (location.pathname !== "/") {
      titleSuffix = location.pathname.split("/")[2];
    }
    //console.log(titleSuffix);
    if (titleSuffix) {
        return (
            <div>
                <Link to="/admin" className='text-decoration-none text-black'>
                <h4>Admin-{titleSuffix}</h4>
                </Link>
                
                <Outlet />
            </div>
        )
    }
    return (
        <div>
            <h2>Admin</h2>
            <ul>
                <li><Link to="create">Create</Link></li>
                <li><Link to="delete">Delete</Link></li>
            </ul>
        </div>
    )
}