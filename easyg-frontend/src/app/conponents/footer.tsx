import Link from "next/link";
export default function Footer() {
    return (
        <div className="mt-5">
            <Link href={"/admin"}>admin</Link>
        </div>
    );
}