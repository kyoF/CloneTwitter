'use client';

import { css } from "@kuma-ui/core";

export const TweetCreateModal = ({ children }: { children: React.ReactNode }) => {
  const containerStyle = css`
    position: fixed;
    z-index: 1;
    left: 0;
    top: 0;
    width: 100%;
    height: 100%;
    overflow: auto;
    background-color: rgba(0, 0, 0, 0.4);
    display: block;
  `;
  const contentStyle = css`
    background-color: #fefefe;
    margin: 15% auto;
    padding: 20px;
    border: 1px solid #888;
    width: 30%;
  `;
  const headerStyle = css`
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 20px;
  `;
  return (
    <div className={containerStyle}>
      <div className={contentStyle}>
        <div className={headerStyle}>
          {children}
        </div>
      </div>
    </div>
  );
}

