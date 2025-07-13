"use client"
import React from 'react';
import Link from 'next/link';
import { logout } from './conponents/logout/logout';
import { useRouter } from 'next/navigation';
export default function Admin() {
    const router = useRouter();

    const handleLogout = async () => {
        await logout(router);
    };

    return (
        <div>
            <ul>
                <li><Link href="/admin/create">Create</Link></li>
                <li><Link href="/admin/delete">Delete</Link></li>
                <li><Link href="#" onClick={handleLogout}>logout</Link></li>
            </ul>
        </div>
    )
}