import { NextRequest } from "next/server"

export async function POST(request:NextRequest) {
  console.log('bbb')
  const body = await request.json()
  console.log(body)
  const res: Response = await fetch(
    `${process.env.API_URL}/tweet/create`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ ...body })
    }
  );
  console.log('aaa');
  const data = await res.json()
  console.log(data);
}
