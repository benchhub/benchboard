# Tensorboard

- Video [Hands-on Tensorboard](https://youtu.be/eBbEDRsCmv4?t=8m22s)
  - video about summary start at 8min, and we only focus on scalar (maybe histogram as well)
  - can use subfolder to identify different runs (i.e. runs/conv=1 runs/conv=2) to compare in graph and filter
  - data seems to be simple protobuf file
    - https://github.com/tensorflow/tensorflow/blob/master/tensorflow/core/framework/summary.proto
- http://dmlc.ml/2017/01/07/bring-TensorBoard-to-MXNet.html
  - the port https://github.com/dmlc/tensorboard to mxnet actually tells more about internal


````protobuf
// Serialization format for histogram module in
// core/lib/histogram/histogram.h
message HistogramProto {
  double min = 1;
  double max = 2;
  double num = 3;
  double sum = 4;
  double sum_squares = 5;

  // Parallel arrays encoding the bucket boundaries and the bucket values.
  // bucket(i) is the count for the bucket i.  The range for
  // a bucket is:
  //   i == 0:  -DBL_MAX .. bucket_limit(0)
  //   i != 0:  bucket_limit(i-1) .. bucket_limit(i)
  repeated double bucket_limit = 6 [packed = true];
  repeated double bucket = 7 [packed = true];
};

// A Summary is a set of named values to be displayed by the
// visualizer.
//
// Summaries are produced regularly during training, as controlled by
// the "summary_interval_secs" attribute of the training operation.
// Summaries are also produced at the end of an evaluation.
message Summary {
  message Image {
    ...
  }

  message Audio {
    ...
  }

  message Value {
    // Tag name for the data. Used by TensorBoard plugins to organize data. Tags
    // are often organized by scope (which contains slashes to convey
    // hierarchy). For example: foo/bar/0
    string tag = 1;

    // Contains metadata on the summary value such as which plugins may use it.
    // Take note that many summary values may lack a metadata field. This is
    // because the FileWriter only keeps a metadata object on the first summary
    // value with a certain tag for each tag. TensorBoard then remembers which
    // tags are associated with which plugins. This saves space.
    SummaryMetadata metadata = 9;

    // Value associated with the tag.
    oneof value {
      float simple_value = 2;
      bytes obsolete_old_style_histogram = 3;
      Image image = 4;
      HistogramProto histo = 5;
      Audio audio = 6;
      TensorProto tensor = 8;
    }
  }

  // Set of values for the summary.
  repeated Value value = 1;
}
````
