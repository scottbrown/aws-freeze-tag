package util

func dir_exists(path) (bool, error) {
  _, err := os.Stat(path)
  if err == nil { return true, nil }
  if os.IsNotExist(err) { return false, nil }
  return true, err
}

func is_dir(path) (bool, error) {
  file_info, err := os.Stat(path)
  if err != nil {
    if os.IsNotExist(err) { return false, nil }
    return false, err
  }
  return file_info.IsDir(), nil
}

func validate_target_dir(target_dir) (bool, error) {
  // ensure dir exists
  exists, err := dir_exists(target_dir)
  if err != nil { return false, err }
  if !exists { return false, nil }

  // ensure dir is a dir
  is_a_dir, err := is_dir(target_dir)
  if err != nil { return false, err }
  if !is_a_dir { return false, nil }

  // ensure user has write access to dir
}

func AWSTagSave() (bool, error) {
  // validate_target_dir()

  instance_id := AWSMetadata("instance-id")
  instance_az := AWSMetadata("placement/availability-zone")
  region := az[0:len(az)-1]

  // tags = retrieve_instance_tags(instance_id, region)
  // map(remove_extraneous_keys, tags)

  // write_tags_to_file(tags, target_dir)
}
