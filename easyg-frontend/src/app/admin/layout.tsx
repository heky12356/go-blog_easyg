import Link from "next/link";

export default async function AdminLayout({ children }: { children: React.ReactNode }) {
    return (
        <div>
            <Link href="/admin" className='text-decoration-none text-black'>
                <h4>Admin</h4>
            </Link>
            {children}
        </div>
    )
}