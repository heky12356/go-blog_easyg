"use client"
import React from 'react';
import Link from 'next/link';
import { logout } from './conponents/logout/logout';
export default function Admin () {
    return (
        <div>
            <ul>
                <li><Link href="/admin/create">Create</Link></li>
                <li><Link href="/admin/delete">Delete</Link></li>
                <li><Link href="#" onClick={logout}>logout</Link></li>
            </ul>
        </div>
    )
}