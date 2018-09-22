package protocol

type BitmapInterface interface {
	GetBitsPerPixel() int
	GetData() [] byte
}


/**
* decompress bitmap from RLE algorithm
* @param	bitmap	{object} bitmap object of bitmap event of node-rdpjs
*/ /*
func decompress (bitmap * pdu.BitmapData) {
	var fName interface {}
	switch (bitmap.bitsPerPixel.value) {
	case 15:
	fName = 'bitmap_decompress_15';
	break;
	case 16:
	fName = 'bitmap_decompress_16';
	break;
	case 24:
	fName = 'bitmap_decompress_24';
	break;
	case 32:
	fName = 'bitmap_decompress_32';
	break;
	default:
	throw 'invalid bitmap data format';
	}

	var input = new Uint8Array(bitmap.bitmapDataStream.value);
	var inputPtr = rle._malloc(input.length);
	var inputHeap = new Uint8Array(rle.HEAPU8.buffer, inputPtr, input.length);
	inputHeap.set(input);

	var ouputSize = bitmap.width.value * bitmap.height.value * 4;
	var outputPtr = rle._malloc(ouputSize);

	var outputHeap = new Uint8Array(rle.HEAPU8.buffer, outputPtr, ouputSize);

	var res = rle.ccall(fName,
	'number',
	['number', 'number', 'number', 'number', 'number', 'number', 'number', 'number'],
	[outputHeap.byteOffset, bitmap.width.value, bitmap.height.value, bitmap.width.value, bitmap.height.value, inputHeap.byteOffset, input.length]
	);

	var output = new Uint8ClampedArray(outputHeap.buffer, outputHeap.byteOffset, ouputSize);

	rle._free(inputPtr);
	rle._free(outputPtr);

	return output;
	} */
