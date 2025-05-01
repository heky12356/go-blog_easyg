export default async function Page({
    params,
  }: {
    params: Promise<{ uid: string }>
  }) {
    const { uid } = await params
    return <div>My Post: {uid}</div>
  }