import React from 'react';
import Link from 'next/link';
export default function Admin () {
    return (
        <div>
            <ul>
                <li><Link href="/admin/create">Create</Link></li>
                <li><Link href="/admin/delete">Delete</Link></li>
            </ul>
        </div>
    )
}