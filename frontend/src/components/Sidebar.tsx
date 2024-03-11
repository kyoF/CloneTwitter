import { css } from "@kuma-ui/core";
import Link from "next/link";

export default function Sidebar() {
  const containerStyle = css`
    width: 250px;
    height: 100vh;
    background-color: #fff;
    border-right: 1px solid #e6ecf0;
    position: fixed;
    left: 0;
    top: 0;
  `;
  const headerStyle = css`
    padding: 10px;
    border-bottom: 1px solid #e6ecf0;
  `;
  const navListStyle = css`
    list-style: none;
    padding: 0;
    margin: 0;
  `;
  const navItemStyle = css`
    padding: 10px;
    cursor: pointer;
    display: flex;
    align-items: center;
    transition: background-color 0.3s ease;
    &:hover {
      background-color: #e8f5fe;
    }
  `;

  return (
    <div className={containerStyle}>
      <div className={headerStyle}>
        <img src="https://logos-world.net/wp-content/uploads/2020/04/Twitter-Logo.png" alt="Twitter Logo" width="30" />
        Twitter
      </div>
      <div className={navListStyle}>
        <div className={navItemStyle}>
          <Link href="/home">home</Link>
        </div>
        <div className={navItemStyle}>
          <Link href="/tweet/create">post</Link>
        </div>
      </div>
    </div>
  );
};

