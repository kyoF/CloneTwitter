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
  return (
    <div className="flex justify-center items-center h-screen bg-gray-200">
      <div className="bg-white rounded-lg shadow-md p-8 max-w-md w-full text-center">
        <img
          src="https://logos-world.net/wp-content/uploads/2020/04/Twitter-Logo.png"
          alt="Twitter Logo"
          width="30"
          className="mr-2"
        />
        <h1 className="text-2xl font-bold mb-6">Log in to Twitter</h1>
        <form>
          <input
            type="text"
            placeholder="Phone, email or username"
            required
            className="mb-4 px-4 py-2 border border-gray-300 rounded-md w-full"
          />
          <input
            type="password"
            placeholder="Password"
            required
            className="mb-6 px-4 py-2 border border-gray-300 rounded-md w-full"
          />
          <button
            type="submit"
            className="bg-twitter text-white font-bold py-2 px-4 rounded-full w-full mb-4"
          >
            Log in
          </button>
        </form>
        <div className="flex justify-center items-center text-sm text-twitter">
          <a href="#" className="hover:underline">
            Forgot password?
          </a>
          <span className="mx-2">&#183;</span>
          <a href="#" className="hover:underline">
            Sign up for Twitter
          </a>
        </div>
      </div>
    </div>
  );
};

