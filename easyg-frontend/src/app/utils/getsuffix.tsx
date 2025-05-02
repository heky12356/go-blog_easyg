import { headers } from 'next/headers';
export default async function Getsuffix(layer : number):Promise<string>{
  const headerList = await headers();
  const pathname = headerList.get('x-current-path') || '/';
  // 解析路径后缀
  const segments = pathname.split('/');
  const titleSuffix = segments[layer] || '';
  return (
    titleSuffix
  );
}