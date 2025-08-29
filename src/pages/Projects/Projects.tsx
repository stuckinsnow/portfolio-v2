import { useApi } from "../../api/useApi";
import type { S3Response } from "./types";

export function Projects() {
  const {
    data: s3Data,
    loading,
    error,
  } = useApi<S3Response>("/api/s3/list-objects");

  const firstImage = s3Data?.objects?.[0];

  return (
    <main>
      <h1>Projects</h1>
      {loading && <p>Loading images...</p>}
      {error && <p>Error: {error}</p>}
      {firstImage ? (
        <div>
          <h2>An S3 Image</h2>
          <img
            src={`/api/s3/object?key=${encodeURIComponent(firstImage)}`}
            alt="First S3 object"
            style={{ maxWidth: "500px", height: "auto" }}
          />
        </div>
      ) : (
        !loading && <p>No images found in S3 bucket</p>
      )}
    </main>
  );
}