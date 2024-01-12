export default async function TweetCreate() {
  const onSubmit = () => {
    console.log("submit");
  }
  return (
    <form onSubmit={onSubmit()}>
      <div>tweet</div>
      <input />
    </form>
  );
}

