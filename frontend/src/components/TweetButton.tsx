export const TweetButton = ({ userId }: { userId: string }) => {
  const tweet = () => {
    console.log(userId);
  };
  return (
    <button onClick={tweet()}>tweet</button>
  );
}
