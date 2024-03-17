import { NextRequest } from "next/server"

export async function POST(request: NextRequest) {
  try {
    const body = await request.json()

    const res: Response = await fetch(
      `${process.env.API_URL}/tweet/create`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ ...body })
    });

    if (!res.ok) {
      throw new Error(`API request failed with status ${res.status}`);
    }

    const data = await res.json();
    return new Response(JSON.stringify(data), {
      status: 200,
      headers: {
        'Content-Type': 'application/json',
      },
    });
  } catch (error) {
    console.error('Error in POST request:', error);
    return new Response(JSON.stringify({ error: 'An error occurred' }), {
      status: 500,
      headers: {
        'Content-Type': 'application/json',
      },
    });
  }
}
