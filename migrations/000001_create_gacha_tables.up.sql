-- -----------------------------------------------------
-- Table `gachas`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `gachas` (
  `id` INT NOT NULL,
  `name` VARCHAR(50) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `name_UNIQUE` (`name` ASC) VISIBLE)
ENGINE = InnoDB;

-- -----------------------------------------------------
-- Table `items`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `items` (
  `id` BIGINT NOT NULL,
  `name` VARCHAR(50) NOT NULL,
  `rarity` VARCHAR(5) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `name_UNIQUE` (`name` ASC) VISIBLE)
ENGINE = InnoDB;

-- -----------------------------------------------------
-- Table `gacha_items`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `gacha_items` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `gacha_id` INT NOT NULL,
  `item_id` BIGINT NOT NULL,
  `weight` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_gacha_items_items_idx` (`item_id` ASC) VISIBLE,
  INDEX `fk_gacha_items_gachas1_idx` (`gacha_id` ASC) VISIBLE,
  CONSTRAINT `fk_gacha_items_items`
    FOREIGN KEY (`item_id`)
    REFERENCES `items` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_gacha_items_gachas1`
    FOREIGN KEY (`gacha_id`)
    REFERENCES `gachas` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;
