export default function Login() {
  const res: Response = await fetch(
    `${process.env.API_URL}/login`,
    { cache: 'no-store' }
  );
  return <div>login</div>;
}

