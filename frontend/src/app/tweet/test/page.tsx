export default async function Fujiki() {
  const userId: string = 'fujiki';
  const postMethod = async (e) => {
    e.preventDefault()

    const res: Response = await fetch(
      `${process.env.API_URL}/user/${userId}/tweet/create`,
      {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(Object.fromEntries(formData)),
      }
    );
  }
  return 
}
