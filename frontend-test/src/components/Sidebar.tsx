import Link from "next/link";

export default function Sidebar() {
  return (
    <div className="w-64 h-screen bg-white border-r border-gray-200 fixed left-0 top-0">
      <div className="px-4 py-2 border-b border-gray-200 flex items-center">
        <img
          src="https://logos-world.net/wp-content/uploads/2020/04/Twitter-Logo.png"
          alt="Twitter Logo"
          width="30"
          className="mr-2"
        />
        Twitter
      </div>
      <nav>
        <ul className="list-none m-0 p-0">
          <li>
            <Link
              href="/home"
              className="block px-4 py-2 flex items-center hover:bg-blue-100 transition-colors duration-300"
            >
              home
            </Link>
          </li>
          <li>
            <Link
              href="/tweet/create"
              className="block px-4 py-2 flex items-center hover:bg-blue-100 transition-colors duration-300"
            >
              post
            </Link>
          </li>
        </ul>
      </nav>
    </div>
  );
}
