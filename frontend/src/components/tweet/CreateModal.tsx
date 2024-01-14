'use client';

import { css } from '@/../styled-system/css'

export const TweetCreateModal = ({ children }: { children: React.ReactNode }) => {
  return (
    <div className={css({
      position: 'fixed',
      top: 0,
      left: 0,
      width: "100%",
      height: "100%",
      backgroundColor: "rgba(0, 0, 0, 0.5)",
      display: 'flex',
      alignItems: 'center',
      justifyContent: 'center'
    })}>
      <div className={css({
        zIndex: 2,
        width: "50%",
        padding: "1em",
        background: "#fff",
      })}>
        {children}
      </div>
    </div>
  );
}

