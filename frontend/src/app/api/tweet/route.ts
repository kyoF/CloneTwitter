import { NextRequest } from "next/server"

export async function POST(request:NextRequest) {
  const body = await request.json()
  const res: Response = await fetch(
    `${process.env.API_URL}/tweet/create`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ ...body })
    }
  );
  const data = await res.json()
}
