export default async function Login() {
  const body = {
    userId: "test",
    email: "test@test.com",
    password: "test-password",
  }
  const res: Response = await fetch(
    `${process.env.API_URL}/login`,
    {
      method: "POST",
      headers: { "Content-Type": "application/json", },
      body: JSON.stringify(body),
      cache: 'no-store'
    }
  );
  return <div>login</div>;
}

