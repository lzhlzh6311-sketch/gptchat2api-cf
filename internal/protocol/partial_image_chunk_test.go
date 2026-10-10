package protocol

import "testing"

func TestProgressChunkForwardsPartialImage(t *testing.T) {
	chunk := ImageOutput{Kind: "progress", Model: "gpt-image-2", Index: 1, Total: 1, UpstreamEventType: "response.image_generation_call.partial_image", PartialImage: "aGVsbG8=", PartialImageIndex: 0, OutputFormat: "png"}.Chunk()
	if chunk["object"] != "image.generation.partial_image" {
		t.Fatalf("object = %v", chunk["object"])
	}
	if chunk["partial_image_b64"] != "aGVsbG8=" || chunk["partial_image_index"] != 0 || chunk["output_format"] != "png" {
		t.Fatalf("unexpected chunk: %#v", chunk)
	}
}

func TestProgressChunkWithoutPartialImageUnchanged(t *testing.T) {
	chunk := ImageOutput{Kind: "progress", Text: "working"}.Chunk()
	if chunk["object"] != "image.generation.chunk" || chunk["progress_text"] != "working" {
		t.Fatalf("unexpected chunk: %#v", chunk)
	}
	if _, ok := chunk["partial_image_b64"]; ok {
		t.Fatal("partial_image_b64 must not be set without a partial image")
	}
}
