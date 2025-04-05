export default async function Page({
    params,
  }: {
    params: Promise<{ test: string }>
  }) {
    const test = (await params).test
    return <div>My Post: {test}</div>
  }