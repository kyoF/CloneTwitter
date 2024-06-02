'use client';

import { useRouter } from 'next/navigation';
import { useState } from 'react';
import { TweetCreateModal } from '@/components/tweet/CreateModal';

export default function TweetCreateModalPage() {
  const router = useRouter();
  const [tweetText, setTweetText] = useState<string>('');
  const [characterCount, setCharacterCount] = useState<number>(0);

  const textChange = (e) => {
    const text = e.target.value;
    setTweetText(text);
    setCharacterCount(text.length);
  };

  const submit = async (e) => {
    e.preventDefault()

    try {
      const res: Response = await fetch('/api/tweet', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ userId: 'fujiki', text: tweetText })
      });

      if (!res.ok) {
        throw new Error(`API request failed with status ${res.status}`);
      }

      setTweetText('');
      setCharacterCount(0);
      router.back();
    } catch (error) {
      console.error('Error submitting tweet:', error);
    }
  };

  const back = () => {
    router.back()
  }

  return (
    <TweetCreateModal>
      <textarea
        placeholder="What's happeing?"
        value={tweetText}
        onChange={textChange}
        maxLength={280}
      />
      <div>
        <span>{characterCount}/280</span>
        <button
          onClick={submit}
        >
          tweet
        </button>
      </div>
      <button onClick={back}>back</button>
    </TweetCreateModal>
  );
}

