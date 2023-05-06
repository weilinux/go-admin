package test

func TestDelete(t *testing.T) {
    user1, user2 := User{Name: "delete1"}, User{Name: "delete2"}
    DB.Save(&user1)
    DB.Save(&user2)

    if err := DB.Delete(&user1).Error; err != nil {
        t.Errorf("No error should happen when delete a record, err=%s", err)
    }

    if !DB.Where("name = ?", user1.Name).First(&User{}).RecordNotFound() {
        t.Errorf("User can't be found after delete")
    }

    if DB.Where("name = ?", user2.Name).First(&User{}).RecordNotFound() {
        t.Errorf("Other users that not deleted should be found-able")
    }
}

func TestInlineDelete(t *testing.T) {
    user1, user2 := User{Name: "inline_delete1"}, User{Name: "inline_delete2"}
    DB.Save(&user1)
    DB.Save(&user2)

    if DB.Delete(&User{}, user1.Id).Error != nil {
        t.Errorf("No error should happen when delete a record")
    } else if !DB.Where("name = ?", user1.Name).First(&User{}).RecordNotFound() {
        t.Errorf("User can't be found after delete")
    }

    if err := DB.Delete(&User{}, "name = ?", user2.Name).Error; err != nil {
        t.Errorf("No error should happen when delete a record, err=%s", err)
    } else if !DB.Where("name = ?", user2.Name).First(&User{}).RecordNotFound() {
        t.Errorf("User can't be found after delete")
    }
}

